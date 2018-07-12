package main
import(
 "fmt"
  "github.com/hashicorp/vault/api"
  //"github.com/hashicorp/consul/connect"
)

func main() {
client, err := api.NewClient(&api.Config{
    Address: "http://127.0.0.1:8200",

})
token:= "e3637bcb-e4d0-b094-33b2-91e7b8f7d59e" //use ur own token
client.SetToken(token)

secretData := map[string]interface{}{
    "value": "world",
    "foo": "bar",
    "age": "-1",
}

secretKey:= "secret/anything"

secretValues, err1 := client.Logical().Read(secretKey)
if err != nil {
   fmt.Println(err1)
}
fmt.Println("secret %s -> %v", secretKey, secretValues)

_, err = client.Logical().Write(secretKey, secretData)
if err!=nil{
	fmt.Println(err)
}
}
