/*
 * Copyright © 2021 ZkBNB Protocol
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package apollo

import (
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/agcache"
	"github.com/apolloconfig/agollo/v4/env/config"
	"testing"
)

func ApolloCient() agcache.CacheInterface {
	c := &config.AppConfig{
		AppID:          "cloud-zkbnb",
		Cluster:        "dev",
		IP:             "http://internal-tf-cb-qa-apollo-config-alb-1199831538.ap-northeast-1.elb.amazonaws.com:9028",
		NamespaceName:  "application",
		IsBackupConfig: true,
	}

	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	fmt.Println("init Apollo config success")
	//Use your apollo key to test
	cache := client.GetConfigCache(c.NamespaceName)
	return cache

}

func Test_Init(t *testing.T) {
	ApolloCient()
}