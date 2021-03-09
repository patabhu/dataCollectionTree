package node

import (
	"dataCollectionTree/models"
	"errors"
	"log"
)

const (
	Webreq    = "webreq"
	Timespent = "timespent"
)

type baseMetric struct {
	WebRequest int
	TimeSpent  int
}
type HeadNode struct {
	baseMetric
	Countries []*CountryNode
}
type CountryNode struct {
	baseMetric
	Country string
	Devices []*DeviceNode
}
type DeviceNode struct {
	Device string
	baseMetric
}

func (n *HeadNode) countryNodes() []*CountryNode {
	return n.Countries
}
func (n *HeadNode) updateWebRequest(webRequest int) {
	n.WebRequest += webRequest
}
func (n *HeadNode) updateTimespent(timeSpent int) {
	n.TimeSpent += timeSpent
}
func (n *CountryNode) updateWebRequest(webRequest int) {
	n.WebRequest += webRequest
}
func (n *CountryNode) updateTimespent(timeSpent int) {
	n.TimeSpent += timeSpent
}
func (n *DeviceNode) updateWebRequest(webRequest int) {
	n.WebRequest += webRequest
}
func (n *DeviceNode) updateTimespent(timeSpent int) {
	n.TimeSpent += timeSpent
}
func newRoot() *HeadNode {
	return new(HeadNode)
}
func newCountry(CountryName string) *CountryNode {
	newCountryNode := new(CountryNode)
	newCountryNode.Country = CountryName
	return newCountryNode
}
func newDevice(deviceName string) *DeviceNode {
	newDeviceNode := new(DeviceNode)
	newDeviceNode.Device = deviceName
	return newDeviceNode
}

// UpdateMetric func
func (head *HeadNode) UpdateMetric(Country string, Device string, metrics *[]models.Metric) (*HeadNode, error) {
	var timespent, webreq int
	for _, m := range *metrics {
		if m.Key == Timespent {
			timespent = m.Value
		}
		if m.Key == Webreq {
			webreq = m.Value
		}
	}
	if head == nil {
		head = newRoot()
	}
	head.updateWebRequest(webreq)
	head.updateTimespent(timespent)

	var countryNode *CountryNode

	//looping through country level children to find the node
	for _, country := range head.countryNodes() {
		if country.Country == Country {
			countryNode = country
			break
		}
	}
	//Create new child node in country level if not present
	if countryNode == nil {
		countryNode = newCountry(Country)
	}

	var deviceNode *DeviceNode

	for _, device := range countryNode.Devices {
		if device.Device == Device {
			deviceNode = device
			break
		}
	}
	if deviceNode == nil {
		deviceNode = newDevice(Device)
	}

	deviceNode.updateTimespent(timespent)
	deviceNode.updateWebRequest(webreq)
	countryNode.Devices = append(countryNode.Devices, deviceNode)

	countryNode.updateTimespent(timespent)
	countryNode.updateWebRequest(webreq)
	head.Countries = append(head.Countries, countryNode)
	return head, nil
}

//GetMetricByCountry gets metrics by country
func (head *HeadNode) GetMetricByCountry(data *models.Data, Country string) error {
	var (
		countryNode *CountryNode
	)
	if head == nil {
		return nil
	}
	//Get the node of the specified country
	// data.Metrics = make(models.Metric{}, 0)
	for _, country := range head.countryNodes() {
		if country.Country == Country {
			countryNode = country
			break
		}
	}

	if countryNode == nil {
		log.Println(errors.New("Country not Found"))
		return errors.New("Country not Found")
	}
	for _, d := range countryNode.Devices {
		log.Println("devices", d)
	}
	//Append the metrics values
	data.Metrics = append(data.Metrics, models.Metric{Key: Webreq, Value: countryNode.WebRequest}, models.Metric{Key: Timespent, Value: countryNode.TimeSpent})
	return nil
}
