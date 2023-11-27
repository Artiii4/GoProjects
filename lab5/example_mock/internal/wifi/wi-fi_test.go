package wifi_test

import (
	"errors"
	basewifi "example_mock/internal/wifi"
	"fmt"
	"net"
	"testing"

	wifi "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

func TestNew(t *testing.T) {
	mockWifi := NewWiFi(t)
	WiFiService := basewifi.Service{WiFi: mockWifi}
	result := basewifi.New(mockWifi)
	if result != WiFiService {
		t.Errorf("expected to be %v, got %v", WiFiService, result)
	}
}

var tableForTestGetAdresses = []struct {
	adress   []string
	rightErr error
}{
	{adress: []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"}, rightErr: nil},
	{adress: []string{"Hello", "World"}, rightErr: errors.New("ExpectedError")},
}

func TestGetAdresses(t *testing.T) {
	mockWifi := NewWiFi(t)
	WifiService := basewifi.Service{WiFi: mockWifi}
	for num, row := range tableForTestGetAdresses {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.adress), row.rightErr)
		realAdress, err := WifiService.GetAddresses()
		if row.rightErr != nil {
			require.ErrorIs(t, err, row.rightErr, "Error!", err)
			continue
		}
		require.NoError(t, err, "row:", num)
		require.Equal(t, parseMACs(row.adress), realAdress,
			"row: %d, right addrs: %s, addrs: %s", num,
			parseMACs(row.adress), realAdress)
	}

}

var tableForTestGetNames = []struct {
	names    []string
	rightErr error
}{
	{names: nil, rightErr: errors.New("ExpectedError")},
	{names: []string{}, rightErr: nil},
}

func TestGetNames(t *testing.T) {
	mockWifi := NewWiFi(t)
	wifiService := basewifi.Service{WiFi: mockWifi}
	for i, row := range tableForTestGetNames {
		mockWifi.On("Interfaces").Unset()
		mockWifi.On("Interfaces").Return(mockIfaces(row.names),
			row.rightErr)
		actualNames, err := wifiService.GetNames()
		if row.rightErr != nil {
			require.ErrorIs(t, err, row.rightErr, "row: %d, right error:  %w, error: %w", i,
				row.rightErr, err)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, actualNames, "row: %d, right address: %s, address: %s", i,
			parseMACs(row.names), actualNames)
	}
}

func mockIfaces(addrs []string) []*wifi.Interface {
	var interfaces []*wifi.Interface
	for i, addrStr := range addrs {
		hwAddr := parseMAC(addrStr)
		if hwAddr == nil {
			continue
		}
		iface := &wifi.Interface{
			Index:        i + 1,
			Name:         fmt.Sprintf("eth%d", i+1),
			HardwareAddr: hwAddr,
			PHY:          1,
			Device:       1,
			Type:         wifi.InterfaceTypeAPVLAN,
			Frequency:    0,
		}
		interfaces = append(interfaces, iface)
	}
	return interfaces
}
func parseMACs(macStr []string) []net.HardwareAddr {
	var addrs []net.HardwareAddr
	for _, addr := range macStr {
		addrs = append(addrs, parseMAC(addr))
	}
	return addrs
}
func parseMAC(macStr string) net.HardwareAddr {
	hwAddr, err := net.ParseMAC(macStr)
	if err != nil {
		return nil
	}
	return hwAddr
}
