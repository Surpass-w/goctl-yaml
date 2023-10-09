package generate

import (
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func Do(filename string, in *plugin.Plugin) error {

	authUnits, err := applyGenerate(in)
	if err != nil {
		fmt.Println(err)
	}

	data, err := yaml.Marshal(authUnits)
	if err != nil {
		fmt.Println(err)
	}

	output := in.Dir + "/" + filename

	err = ioutil.WriteFile(output, data, 0666)
	if err != nil {
		fmt.Println(err)
	}

	return err
}
