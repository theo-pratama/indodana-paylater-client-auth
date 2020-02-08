package main

import "crypto/hmac"
import "crypto/sha256"
import "os"
import "time"
import "strconv"
import "fmt"

func main() {
 nonce := time.Now().UnixNano() / 1e6 / 1000
 apiKey := os.Getenv("API_KEY")
 apiSecret := os.Getenv("API_SECRET")

 if (apiKey == "") {
  fmt.Printf("Api Key :\n");
  fmt.Scanf("%s", &apiKey);
 }


 if (apiSecret == "") {
  fmt.Printf("Api Secret :\n");
  fmt.Scanf("%s", &apiSecret);
 }

 nonceStr := strconv.Itoa(int(nonce))
 content := []byte(apiKey + ":" + nonceStr)

 mac := hmac.New(sha256.New, []byte(apiSecret))
 mac.Write(content)
 result := mac.Sum(nil)

 fmt.Printf("Bearer %s:%s:%x\n", apiKey, nonceStr, result)
}
