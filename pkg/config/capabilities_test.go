/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"testing"

	cb "github.com/hyperledger/fabric-protos-go/common"
	. "github.com/onsi/gomega"
)

func TestGetChannelCapabilities(t *testing.T) {
	t.Parallel()

	gt := NewGomegaWithT(t)

	expectedCapabilities := map[string]bool{"V1_3": true}

	config := &cb.Config{
		ChannelGroup: &cb.ConfigGroup{
			Values: map[string]*cb.ConfigValue{},
		},
	}

	c := ConfigTx{
		base:    config,
		updated: config,
	}

	err := addValue(config.ChannelGroup, capabilitiesValue(expectedCapabilities), AdminsPolicyKey)
	gt.Expect(err).NotTo(HaveOccurred())

	channelCapabilities, err := c.GetChannelCapabilities()
	gt.Expect(err).NotTo(HaveOccurred())
	gt.Expect(channelCapabilities).To(Equal(expectedCapabilities))

	// Delete the capabilities key and assert retrieval to return nil
	delete(config.ChannelGroup.Values, CapabilitiesKey)
	channelCapabilities, err = c.GetChannelCapabilities()
	gt.Expect(err).NotTo(HaveOccurred())
	gt.Expect(channelCapabilities).To(BeNil())
}

func TestGetOrdererCapabilities(t *testing.T) {
	t.Parallel()

	gt := NewGomegaWithT(t)

	baseOrdererConf := baseSoloOrderer(t)
	ordererGroup, err := newOrdererGroup(baseOrdererConf)
	gt.Expect(err).NotTo(HaveOccurred())

	config := &cb.Config{
		ChannelGroup: &cb.ConfigGroup{
			Groups: map[string]*cb.ConfigGroup{
				OrdererGroupKey: ordererGroup,
			},
		},
	}

	c := ConfigTx{
		base:    config,
		updated: config,
	}

	ordererCapabilities, err := c.GetOrdererCapabilities()
	gt.Expect(err).NotTo(HaveOccurred())
	gt.Expect(ordererCapabilities).To(Equal(baseOrdererConf.Capabilities))

	// Delete the capabilities key and assert retrieval to return nil
	delete(config.ChannelGroup.Groups[OrdererGroupKey].Values, CapabilitiesKey)
	ordererCapabilities, err = c.GetOrdererCapabilities()
	gt.Expect(err).NotTo(HaveOccurred())
	gt.Expect(ordererCapabilities).To(BeNil())
}

func TestGetOrdererCapabilitiesFailure(t *testing.T) {
	t.Parallel()

	gt := NewGomegaWithT(t)

	config := &cb.Config{
		ChannelGroup: &cb.ConfigGroup{
			Groups: map[string]*cb.ConfigGroup{},
		},
	}

	c := ConfigTx{
		base:    config,
		updated: config,
	}

	ordererCapabilities, err := c.GetOrdererCapabilities()
	gt.Expect(err).To(MatchError("orderer missing from config"))
	gt.Expect(ordererCapabilities).To(BeNil())
}

func TestGetApplicationCapabilities(t *testing.T) {
	t.Parallel()

	gt := NewGomegaWithT(t)

	baseApplicationConf := baseApplication(t)
	applicationGroup, err := newApplicationGroup(baseApplicationConf)
	gt.Expect(err).NotTo(HaveOccurred())

	config := &cb.Config{
		ChannelGroup: &cb.ConfigGroup{
			Groups: map[string]*cb.ConfigGroup{
				ApplicationGroupKey: applicationGroup,
			},
		},
	}

	c := ConfigTx{
		base:    config,
		updated: config,
	}

	applicationCapabilities, err := c.GetApplicationCapabilities()
	gt.Expect(err).NotTo(HaveOccurred())
	gt.Expect(applicationCapabilities).To(Equal(baseApplicationConf.Capabilities))

	// Delete the capabilities key and assert retrieval to return nil
	delete(config.ChannelGroup.Groups[ApplicationGroupKey].Values, CapabilitiesKey)
	applicationCapabilities, err = c.GetApplicationCapabilities()
	gt.Expect(err).NotTo(HaveOccurred())
	gt.Expect(applicationCapabilities).To(BeNil())
}

func TestGetApplicationCapabilitiesFailure(t *testing.T) {
	t.Parallel()

	gt := NewGomegaWithT(t)

	config := &cb.Config{
		ChannelGroup: &cb.ConfigGroup{
			Groups: map[string]*cb.ConfigGroup{},
		},
	}

	c := ConfigTx{
		base:    config,
		updated: config,
	}

	applicationCapabilities, err := c.GetApplicationCapabilities()
	gt.Expect(err).To(MatchError("application missing from config"))
	gt.Expect(applicationCapabilities).To(BeNil())
}
