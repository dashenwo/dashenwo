package gorm

import (
	conf "github.com/dashenwo/dashenwo/console/account/config"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/util/log"
	"github.com/olivere/elastic/v7"
	"sync"
)

var (
	esonce   sync.Once
	esConf   conf.Elasticsearch
	esClient *elastic.Client
)

func InitElasticsearch() {
	esonce.Do(func() {
		esConf = conf.Elasticsearch{}
		err := config.Get("elasticsearch").Scan(&esConf)
		if err != nil {
			log.Fatal(err)
		}
		var options []elastic.ClientOptionFunc
		// 配置host
		if esConf.Hosts != nil {
			options = append(options, elastic.SetURL(esConf.Hosts...))
		}
		// 如果配置了不检测地址，在调试的时候可用
		if esConf.Sniff == false {
			options = append(options, elastic.SetSniff(esConf.Sniff))
		}
		// 如果有用户名和密码
		if esConf.BasicAuth.UserName != "" {
			elastic.SetBasicAuth(esConf.BasicAuth.UserName, esConf.BasicAuth.PassWord)
		}
		// 创建一个es客户端
		esClient, err = elastic.NewClient(options...)
		if err != nil {
			log.Fatal(err)
		}
	})
}
