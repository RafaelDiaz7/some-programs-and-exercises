package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var demoText = "nananananan bbbbbbbbb...___7863,,,,'''"
	//var demoText = "Lorem ipsum dolor sit amet consectetur adipiscing elit nisi, condimentum blandit elementum bibendum sed potenti accumsan faucibus ligula, euismod netus justo auctor luctus tincidunt gravida. At etiam litora massa parturient quis blandit class ridiculus taciti auctor molestie himenaeos potenti, a in neque dui praesent enim sem risus ac viverra dictumst cum, cubilia senectus et placerat vel quisque lacinia accumsan natoque aliquam semper tempus. Dapibus egestas elementum aliquet aliquam parturient euismod eu vulputate nulla magnis, non semper nam tempus commodo quam fames inceptos vitae, purus sapien augue suspendisse id hendrerit cum fermentum luctus.\n\nFringilla primis lectus sociis habitant tincidunt potenti tempus conubia, libero vivamus dictumst aliquet condimentum hac magnis, posuere viverra massa imperdiet nulla et purus. Viverra cursus mi lectus neque enim nisi pretium suscipit dis, scelerisque malesuada sollicitudin venenatis facilisi ultrices nulla massa. Integer platea hac vehicula commodo phasellus a varius, duis curabitur cras sociosqu nunc mollis torquent natoque, luctus euismod primis litora non lobortis."
	demoText = strings.ReplaceAll(demoText, ".", "")
	fmt.Println(demoText)
	normalizedText := normalize(demoText)
	fmt.Println(normalizedText)
	fmt.Println(countLetters(normalizedText))
}

func normalize(text string) string {
	re := regexp.MustCompile(`[_?!.,')(/;:1234567890 ]`)
	return strings.ToLower(re.ReplaceAllString(text, ``))
}

func countLetters(text string) map[string]uint8 {
	var letters = make(map[string]uint8)
	for i := 0; i <= len(text)-1; i++ {
			letters[string(text[i])]++
	}
	return letters
}
