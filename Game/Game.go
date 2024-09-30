package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var Finish bool //variable who indicate if the game if finish or not.
var Win bool    //variable to indicate if the user win or not

func Game() {
	/*
		Function who are in charge of the game.
	*/
	rand.Seed(time.Now().UnixNano())
	Tries = 12
	Finish = false
	Win = false
	RandomLetter := rune(Word[rand.Intn(Lenght)])
	ListToDisplay := []rune{}
	AlreadyTry := []string{}
	AlreadyTryCorrect := []rune{RandomLetter}
	LetterFind := 0
	for i := 0; i < Lenght; i++ {
		if rune(Word[i]) == RandomLetter {
			ListToDisplay = append(ListToDisplay, RandomLetter)
			LetterFind++
		} else {
			ListToDisplay = append(ListToDisplay, '_')
		}
	}
	for !Finish {
		ClearScreen()
		fmt.Println(Red, " ============================== Hangman game ==============================")
		DisplayNbGameDifficulty()
		temp := ""
		for _, element := range ListToDisplay {
			temp += string(element) + " "
		}
		temp2 := ""
		for i := 0; i < 31-len(temp); i++ {
			temp2 += " "
		}
		fmt.Println(Yellow, "|", Violet, "                       Word to Find :", Green, temp, Yellow, temp2+"|")
		DisplayAlreadyTry(AlreadyTry)
		if Tries >= 10 {
			temp2 = ""
		} else {
			temp2 = " "
		}
		fmt.Println(Yellow, "|", Red, "                        Tries remaining :", Cyan, Tries, Yellow, temp2+"                         |")
		DisplayHangman(Tries)
		fmt.Println(Red, " ==========================================================================")
		fmt.Printf("%s\n\nEnter a letter or a word : ", White)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		temp = scanner.Text()
		temp = strings.ToLower(temp)
		temp = strings.ReplaceAll(temp, " ", "")
		if len(temp) == 1 {
			SubmitesLetter := rune(temp[0])
			if SubmitesLetter >= 'a' && SubmitesLetter <= 'z' {
				FindInWord := false
				for c := 0; c < Lenght; c++ {
					if SubmitesLetter == rune(Word[c]) && ListToDisplay[c] == '_' {
						ListToDisplay[c] = rune(SubmitesLetter)
						FindInWord = true
						LetterFind++
						AlreadyTryCorrect = append(AlreadyTryCorrect, rune(SubmitesLetter))
					}
				}
				for _, element := range AlreadyTryCorrect {
					if SubmitesLetter == element {
						FindInWord = true
					}
				}
				for _, element := range AlreadyTry {
					if len(element) == 1 {
						if SubmitesLetter == rune(element[0]) {
							FindInWord = true
						}
					}
				}
				if !FindInWord {
					AlreadyTry = append(AlreadyTry, temp)
					Tries--
				}
			}
		} else {
			if !strings.ContainsAny(temp, "\n\t\r!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ 0123456789 àâäéèêëîïôöùûüÿç±÷×√∑∏∫∞≈≠≤≥∂∆∇∴∵∪∩∧∨∅∈∉∀∃∝∠⊥⊢⊨⊕⊗πφ°©®€£¥¤§¶†‡•‰′″¹²³½¼¾‖¬¿¡‹›«»▓Ì»▀░¦ß▬▬▬!¢Ø║Ö╣õùÖậ�¦Æ×ØìѰѱѹҘҙҫӂӁѷӑӕӗӚәӛӝӟӞӜӢӣӤӥӦӧӨөӪӫӬӭӮӯӰӱӲӳӴӵӸӹԁẠạẢảẤấẦầẨẩẪẫẬậẮắẰằẲẳẴẵẶặẸẹẺẻẼẽẾếỀềỂểỄễỆệỈỉỊịỌọ😀😁😂🤣😃😄😅😆😉😊😋😎😍😘🥰😗😙😚🙂🤗🤩🤔🤨😐😑😶🙄😏😣😥😮🤐😯😪😫😴😌😛😜😝🤤😒😓😔😕🙃🤑😲☹🙁😖😞😟😤😢😭😦😧😨😩🤯😬😰😱🥵🥶😳🤪😵😡😠🤬😷🤒🤕🤢🤮🤧😇🥳🥴🥺🤠🤡🤥🤫🤭🧐🤓😈👿👹👺💀👻👽👾🤖💩😺😸😹😻😼😽🙀😿😾👶🧒👦👧🧑👨👩🧓👴👵👨‍⚕️👩‍⚕️👨‍🎓👩‍🎓👨‍🏫👩‍🏫👨‍⚖️👩‍⚖️👨‍🌾👩‍🌾👨‍🍳👩‍🍳👨‍🔧👩‍🔧👨‍🏭👩‍🏭👨‍💼👩‍💼👨‍🔬👩‍🔬👨‍💻👩‍💻👨‍🎤👩‍🎤👨‍🎨👩‍🎨👨‍✈️👩‍✈️👨‍🚀👩‍🚀👨‍🚒👩‍🚒👮‍♂️👮‍♀️👷‍♂️👷‍♀️💂‍♂️💂‍♀️🕵️‍♂️🕵️‍♀️👳‍♂️👳‍♀️👲👱‍♂️👱‍♀️🤵‍♂️🤵‍♀️👰‍♂️👰‍♀️🤰🤱👼🎅🤶🦸‍♂️🦸‍♀️🦹‍♂️🦹‍♀️🧙‍♂️🧙‍♀️🧚‍♂️🧚‍♀️🧛‍♂️🧛‍♀️🧜‍♂️🧜‍♀️🧝‍♂️🧝‍♀️🧞‍♂️🧞‍♀️🧟‍♂️🧟‍♀️💆‍♂️💆‍♀️💇‍♂️💇‍♀️🚶‍♂️🚶‍♀️🏃‍♂️🏃‍♀️💃🕺🕴️👯‍♂️👯‍♀️🧖‍♂️🧖‍♀️🧘‍♂️🧘‍♀️🛀🛌👫👭👬💏💑👪👨‍👩‍👦👨‍👩‍👧👨‍👩‍👧‍👦👨‍👨‍👦👨‍👨‍👧👨‍👨‍👧‍👦👨‍👩‍👦‍👦👩‍👩‍👦‍👦👩‍👩‍👧‍👧👩‍👩‍👧🧶🧵🛍️🎒🧳👓🕶️🥽🥼👔👕👖🧣🧤🧥🧦👗👘🥻🩱🩳👙👚👛👜👝🎩🧢👑👒🎓⛑️👢👞👟🥾🥿👠👡🩰👑👒📿💄💍💎🐶🐱🐭🐹🐰🦊🐻🐼🐨🐯🦁🐮🐷🐽🐸🐵🙈🙉🙊🐒🐔🐧🐦🐤🐣🐥🦆🦅🦉🦇🐺🐗🐴🦄🐝🪲🐞🐜🦟🦠🐛🦋🐌🐚🐞🐢🐍🦎🦂🦀🦞🦐🐡🐠🐟🐬🐳🐋🦈🐊🦭🦧🦮🦥🦦🦙🐃🐂🐄🐎🐑🐐🦌🐕🐩🐈🐈‍⬛🐅🐆🦓🦍🦧🐘🦏🐪🐫🦘🦏🐁🐀🐇🐉🦕🦖🐾🐾🐲🐉🔥🌪️🌈☀️🌙🌎🌍🌏💫⭐🌟✨⚡❄️☃️☄️🌧️⛈️🌦️🌥️🌤️⛅🌪️🌈🌋🌏🌍🌎🗻⛰️🏔️🏕️🏖️🏜️🏝️🏞️🗺️🌐🗾🏞️가각간갇갈갉갊감갑값갓갔강갖갗같갚갛개객갠갤갬갭갯갰갱갸갹갼걀걋걍거걱건걷걸걺검겁것겄겅게겐겉겊겋겍겐겔겜겝겟겠겡겨격견결겹겻경계고곡곤곧골곪곬곰곱곳공과관괄광괘괴굉교구국군굳굴굵굶굼굽궁권궐궜귀규균귤그극근글긁금급긋긍기긴길김깅깊까깍깎깐깔깜깝깟깠깡깨깬깰깸깹깻깼깽꼐꼬꼭꼰꼴꼼꼽꽁꽂꽃꽉꽤꽥꽹꾀꾄꾐꾑꾕꾜꾸꾹꾼꿀꿇꿈꿉꿋꿎꿍뀌뀜뀝뀨끄끅끈끊끌끎끓끔끕끗끙끼끽낀낄낌낍낏낐낑的一是在不了有和人这中大为上个国我以要他时来用们生到作地于出就分对成会可主发年动同工也能下过子说产种面而方后多定行学法所民得经十三之到民生见时然子生经而使ا ب ت ث ج ح خ د ذ ر ز س ش ص ض ط ظ ع غ ف ق ك ل م ن ه و يあいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをんアイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲンАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯกขฃคฅฆงจฉชซฌญฎฏฐฑฒณดตถทธนบปผฝพฟภมยรฤลฦวศษสหฬอฮა ბ გ დ ე ვ ზ თ ი კ ლ მ ნ ო პ ჟ რ ს ტ უ ფ ქ ღ ყ შ ჩ ც ძ წ ხ ჯ ჰԱ Բ Գ Դ Ե Զ Է Ը Թ Ժ Ի Լ Խ Ծ Կ Հ Ձ Ղ Ճ Մ Յ Ն Շ Ո Չ Պ Ջ Ռ Ս Վ Տ Ր Ց Ու Փ Ք Օ Ֆअ आ इ ई उ ऊ ऋ ए ऐ ओ औ क ख ग घ ङ च छ ज झ ञ ट ठ ड ढ ण त थ द ध न प फ ब भ म य र ल व श ष स हཀ ཁ ག ང ཅ ཆ ཇ ཉ ཏ ཐ ད ན པ ཕ བ མ ཙ ཚ ཛ ཝ ཞ ཟ འ ཡ ར ལ ཤ ས ཧ ཨঅ আ ই ঈ উ ঊ ঋ এ ঐ ও ঔ ক খ গ ঘ ঙ চ ছ জ ঝ ঞ ট ঠ ড ঢ ণ ত থ দ ধ ন প ফ ব ভ ম য র ল শ ষ স হក ខ គ ឃ ង ច ឆ ជ ឈ ញ ដ ឋ ឌ ឍ ណ ត ថ ទ ធ ន ប ផ ព ភ ម យ រ ល វ ស ហ ឡ អက ခ ဂ ဃ င စ ဆ ဇ ဈ ဉ ဋ ဌ ဍ ဎ ဏ တ ထ ဒ ဓ န ပ ဖ ဗ ဘ မ ယ ရ လ ဝ သ ဟ ဠ အΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ") && len(temp) < 46 {
				if temp == Word {
					LetterFind = Lenght
				} else {
					IsAlreadyTry := false
					for _, element := range AlreadyTry {
						if len(element) != 1 && temp == element {
							IsAlreadyTry = true
						}
					}
					if !IsAlreadyTry {
						AlreadyTry = append(AlreadyTry, temp)
						Tries -= 2
					}
				}
			}
		}
		if LetterFind == Lenght {
			Finish = true
			Win = true
		} else {
			if Tries <= 0 {
				Finish = true
			}
		}
	}
}
