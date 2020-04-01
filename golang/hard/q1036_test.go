package hard

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	var blocked [][]int
	var source []int
	var target []int
	blocked = [][]int{{859360, 747558}, {767839, 414499}, {957307, 53074}, {680104, 950973}, {777486, 371611}, {400205, 636022}, {480897, 491725}, {599558, 635067}, {604926, 321303}, {146897, 874078}, {69813, 369743}, {240303, 869910}, {297688, 930476}, {819428, 899389}, {528032, 699402}, {442, 726800}, {42886, 994090}, {222527, 521523}, {227403, 25583}, {542274, 603306}, {428336, 730145}, {647514, 773345}, {783807, 944686}, {284767, 243365}, {191104, 475620}, {133935, 672361}, {751876, 609273}, {162090, 772994}, {953887, 863427}, {918711, 683947}, {68517, 745792}, {281213, 901250}, {649332, 565904}, {211498, 873161}, {472298, 576437}, {657648, 423014}, {322155, 223386}, {104451, 548963}, {832668, 155145}, {310940, 588911}, {102891, 729651}, {893369, 81819}, {60100, 810221}, {586906, 722845}, {184989, 88941}, {312646, 650448}, {621355, 557009}, {787268, 773811}, {978155, 882776}, {775843, 758630}, {63834, 988711}, {345650, 156464}, {890173, 358341}, {640840, 107643}, {155546, 506015}, {415037, 470166}, {783677, 827973}, {70511, 424498}, {726709, 165867}, {315711, 921062}, {38804, 847785}, {295889, 787820}, {953842, 224828}, {768592, 129889}, {283946, 933122}, {26535, 656712}, {98331, 673090}, {389065, 728072}, {676328, 741858}, {626345, 724074}, {992313, 344431}, {759811, 831836}, {700408, 665625}, {877522, 468336}, {292117, 463603}, {74647, 814621}, {707459, 101239}, {621145, 461629}, {750776, 755457}, {413855, 870990}, {680497, 421583}}
	source = []int{182699, 459911}
	target = []int{925304, 621233}
	r := isEscapePossible(blocked, source, target)
	fmt.Printf("%v\n", r)
}