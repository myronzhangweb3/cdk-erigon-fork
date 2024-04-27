package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 原始JSON结构
type Original struct {
	L1Config struct {
		ChainId                           int    `json:"chainId"`
		CdkValidiumAddress                string `json:"cdkValidiumAddress"`
		MaticTokenAddress                 string `json:"maticTokenAddress"`
		PolygonZkEVMGlobalExitRootAddress string `json:"polygonZkEVMGlobalExitRootAddress"`
		CdkDataCommitteeContract          string `json:"cdkDataCommitteeContract"`
	} `json:"l1Config"`
	GenesisBlockNumber int    `json:"genesisBlockNumber"`
	Root               string `json:"root"`
	Genesis            []struct {
		ContractName string            `json:"contractName"`
		AccountName  string            `json:"accountName"`
		Balance      string            `json:"balance"`
		Nonce        string            `json:"nonce"`
		Address      string            `json:"address"`
		Bytecode     string            `json:"bytecode"`
		Storage      map[string]string `json:"storage"`
	} `json:"genesis"`
}

type OutputAccount struct {
	Balance string      `json:"balance"`
	Nonce   string      `json:"nonce"`
	Code    interface{} `json:"code"`
	Storage interface{} `json:"storage"`

	AccountName string `json:"accountName"`
}

type OutputContract struct {
	Balance string      `json:"balance"`
	Nonce   string      `json:"nonce"`
	Code    interface{} `json:"code"`
	Storage interface{} `json:"storage"`

	ContractName string `json:"contractName"`
}

func main() {
	// 读取文件
	filePath := "genesis_merlin.json" // 更改为实际文件路径
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// 解析原始JSON
	var original Original
	err = json.Unmarshal(file, &original)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	// 转换数据
	transformed := make(map[string]interface{}, 10)
	for _, g := range original.Genesis {
		if g.Bytecode == "" {
			transformed[g.Address] = OutputAccount{
				AccountName: g.AccountName,
				Balance:     g.Balance,
				Nonce:       g.Nonce,
				Code:        nil,
				Storage:     g.Storage,
			}
		} else {
			transformed[g.Address] = OutputContract{
				ContractName: g.ContractName,
				Balance:      g.Balance,
				Nonce:        g.Nonce,
				Code:         g.Bytecode,
				Storage:      g.Storage,
			}
		}
	}

	// 编码为JSON
	result, err := json.MarshalIndent(transformed, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		os.Exit(1)
	}

	// 输出转换后的JSON
	fmt.Println(string(result))
}
