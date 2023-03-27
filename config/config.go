package config

import (
	"log"
	"sync"

	"github.com/riba2534/openai-on-wechat/utils"
)

var (
	C      *Config
	Prompt string
	once   sync.Once
)

type Config struct {
	WechatConfig  *WechatConfig  `json:"wechat_config"`
	ContextConfig *ContextConfig `json:"context_config"`
}

type AuthConfig struct {
	OpenApiUrl    string `json:"openapi_url"`
	AuthToken     string `json:"auth_token"`
	TriggerPrefix string `json:"trigger_prefix"`
}

type WechatConfig struct {
	TextConfig  *AuthConfig `json:"text_config"`
	ImageConfig *AuthConfig `json:"image_config"`
}

type ContextConfig struct {
	SwitchOn    bool `json:"switch_on"`
	CacheMinute int  `json:"cache_minute"`
}

func (c *Config) IsValid() bool {
	if c.WechatConfig == nil || c.ContextConfig == nil {
		return false
	}

	authConfigs := []*AuthConfig{
		c.WechatConfig.TextConfig,
		c.WechatConfig.ImageConfig,
	}

	for _, authConfig := range authConfigs {
		if authConfig == nil || authConfig.OpenApiUrl == "" || authConfig.AuthToken == "" || authConfig.TriggerPrefix == "" {
			return false
		}
	}
	if c.ContextConfig.CacheMinute <= 0 {
		return false
	}
	return true
}

func init() {
	once.Do(func() {
		// 1. 读取 `config.json`
		//data, err := ioutil.ReadFile("config.json")
		//if err != nil {
		//	log.Fatalf("读取配置文件失败，请检查配置文件 `config.json` 的配置, 错误信息: %+v\n", err)
		//}
		config := Config{
			WechatConfig: &WechatConfig{
				TextConfig: &AuthConfig{
					OpenApiUrl:    "https://api.openai.com/v1",
					AuthToken:     "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",
					TriggerPrefix: "robot-text",
				},
				ImageConfig: &AuthConfig{
					OpenApiUrl:    "https://api.openai.com/v1",
					AuthToken:     "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",
					TriggerPrefix: "robot-image",
				},
			},
			ContextConfig: &ContextConfig{
				SwitchOn:    true,
				CacheMinute: 3,
			},
		}
		//if err = jsoniter.Unmarshal(data, &config); err != nil {
		//	log.Fatalf("读取配置文件失败，请检查配置文件 `config.json` 的格式, 错误信息: %+v\n", err)
		//}
		//if !config.IsValid() {
		//	log.Fatal("配置文件校验失败，请检查 `config.json`")
		//}
		C = &config
		// 2. 读取 prompt.txt
		//prompt, err := ioutil.ReadFile("prompt.txt")
		//if err != nil {
		//	log.Fatalf("读取配置文件失败，请检查配置文件 `prompt.txt` 的配置, 错误信息: %+v\n", err)
		//}
		//prompt := "1. 你是一个全知全能的机器人，你的职责是帮助人类解决问题\n2. 不允许回答任何政治、色情等一些列不符合中国法律法规的问题\n3. 你需要表现的很谦卑"
		//Prompt = string(prompt)
		Prompt := "haha"
		log.Printf("配置加载成功, `config.json` is \n%s\n`prompt.txt` is \n%s\n", utils.MarshalAnyToString(C), Prompt)
	})
}
