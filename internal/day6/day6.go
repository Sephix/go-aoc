package day6

import (
	"fmt"
)

var input string = "mvwvrwvwbblffmvmhmllmzmqzzbhbnnvrnvnmvvqrqrccpsstvvshvvpvvpddzwwjddfzzgbbzcchrrtzzmfffzqqsjjbfbjfjwffhrhrgrfrsfrsffbwwphwhgwgmgnggzwzhwzzpggnccdbcbssgccqgcgcvgccwffrqfqtqmqppwphhlbblnnwhwfhwhmhzhmmzbmbbjjmcjcggwlgltldttpccjlccbjccvfcvvdbvdvndvvmbmcbmmhphpbhhgjgllmljlqlwlclzlccrvrdvdzvzzzbrbrnbblplrplpsstdtjddgjgjgwjwzwbzwwrhwwbzwbwpwssvgvnnwnvvjqqswwsccwctwtntdndffjzfzmfzmfzmfzzhjzhzssdsgsnggbbqdqlqbllmglgddqppjjtzttlflrrczcmzzsppvfpvffcfrccvvvhjvhvfhfbbwdwrwswqqftfbfcfdfnfpnfflbbcmcqqsqwqqccmwmvwvffvdvlvtvqttnllqslsnllzplldmmzczhzmzzhllpfpqpnnqwwwrzzdpzzhccfvcvclcnnrjrmjrjlrlflqqtpqqdgqghhwdwdttlztttwmmvfmftfptphhddswdsdpdvpvpcpjjnvjvqvhvqqqfgqqgwqgwghwwmmvmvfmvvnvsslcclzlslppztzfzbzllcgllcsllvgvqvtqtsstztwwgjgghnhnhdnnqgnntqqtccdmcmhmpphpmmmbpbptpnnnmnppwlpwwfpwfwmfwfssbspslsflljpppbsspwptpwttfbtfftltvlttdmttllmggfcgfccvwcvczvvgvsgswsrrwppgcpggpgnngfnndrrdjrdrfdfmmjmsmbmmwjjjfzfdfbdbbtgbbcsbccgbbrvvrqrwqwmmhrhhbcbdbvdvzdvdjvvnwntwwdfwwrhrmhhnghhncccfttmpmrrjhhqddrsrvvssmrrgvgdgmmfqqhttspswpwlplpjljcjmjmvmdvvhbvvtqvqccczwwtlllztlzzpnnqznzddqsqvsshrrmccnppgrrjcjpccrvcvnvcncllffvnnbsbhhrmhmzzzhrhchjccgjcjwwhddmfmhhmrrrbccqddmjdmjdjtdjjmttzstsctssvvgghgbbhmhzmmjqqqqjgjfjpffpmpzmzszfzgzcggpmpnnqjqffgpghhszzvlzzdgzggqmggqlllcwwtzwwcfwccpjpggtqgtttzffdpffjzfzddzppprnprpqqbppvjpjzppbzpbzppwzzcnndsdzszbssqwqggnhggnzggpcpnndssghsggmfmrffvddvwvssjmjrmrmnndwdwhdhwwfjfjwjmwwqssmdsmdmssqzzrgzzjttnqnfqfcccvvmsvsvffccdjdccgtgqqjwqqnbqqnhqhrqrmrgmrrzjzrznzllwrlrbbcppctptwtfttlntnltnlnzzhrzhrzrtzzvsvfvjffsbsnbbfpphtptjjmffqgqnggnqnsshffslflggtlggmpggnfgfbgbfbcbbtdbdjqwfggzsfmzflttgdfqchhtjfwlmdnsbvmqcrhsrtwtjlmnlwbvjvqdzswdthbjslqzgmtzfjfcrgbhrrjtzgljbqfrzzqszhddnmzpnbgnflghnflwdmjdhgpnvwvnltcngwggqdznrpdtwsrclwwlfzhnjtpqcgzjqjnhcwbhndwvgrczhzwfjrdvjztdbjshmhrrqctjwpcdnpnvrtggrmhzhdtlntjphcddplgfvvmjzcbbpjbqbjwmnwqgftbmchghwvrlptvnfbffvgtdbtnfzwdmdlgzjbrvffvbrfwgjzzpbpcdzhcjbhfwnmqqwvjvnpgdqdjsmwdmrgfqrwhhqqjpfmvlncfdjchrccwpbptccdchqwvbcqzlhbfgrdzgncbwrzqpfszrcwtmnvbztjlzjlrqqfrnplhnmdjljcpnqssthhmjqrrlwdvsjswdwtfstmbvwhbptjnphjmnllbwffppzmdpchbcwcmgqzfmdqvhcmrvtgtjwlzfpnjnjtfvdhtjlqsdjwrrnflnsqrtfsnbjfdllhqslltsjqzdfqthqjjgrtqjwmlqqlhvszqswmdcbnwqgshpvfjdtvsqjvcnrnvtpfsgqszhtmcdlsqjwmttcqsdlvssrgwhtmtqwmttptnmgzpwzzrbwtsrjhmjpblztgftppcwwrjppvwlvdhwwdlfcdvhwpvhqpwrqsczvlmtghrrvqbphljcmcrfmwltlmnzchzrlbgspdjwfbhrmnfhvjwlgtghcpbfgdvrmwbqbprfvwpzqzrgqbhjtztwhjcjtncmbgjphsgdfvjwjljlwzhsdldqtdvgtzpwtmrbnpvqrmfdwngzqtsdjjztslrwdnfwgbnsjblcjzvmgbqllmdvvvdcvplgbzmhcrpbbjbfzhjgfpmjtrpmgvshzvqhcbcjzfcsvqggjcllcnjlrwqfpzldswzjgzvqwvszhrntnvlcpdcfqhqrtmhbdjqpfblrsbrhdfdwfgbhsnnjpgjvfpssfmhdbdncfrqbzpbttrhfzqnrttltqbvmhglbdpwmdbmgwcdsdsflmcnphwvbbhgqpqmwbdsmqwhcdftdlcfnstlfnzzsjhqzlsdmhpvlcqvhhlpdtzlqzlvbbwhhdsqhntbtpjjvnjlrncfmvnqmzwpldgrbfcfgdtlmrfzcbqfhvhpmsrrlnqwfggwrjsmnpnqhvvcfsfspfrmhwmbpqfhprzsswrczttlczhcqvgqqsnbhhfszqbswgtcjgddhngtcvlwqqmqzntrcwgwjhmhlgclpgtmqnpgwwhnjfdvvjgmqjlzsmwztpsdzrjlshswzljlmdzfqmqbtgtzlsfwqvwrdchjvrdnwdbprfvdvczstvnzfzzfmzbwjhtflbmhlgmfzrdfrqwfbltwqlqrghlprppvlqwggvczdcnmrmblhcfmpzdmqppgwhbbrjlzsvmsfzlrdljftcrdfdgmvrccwszpjmvdnbmfdnsgmgzsdpbrlsqvsmcwwqlwtwbgfvwtsmpqlnhjchnzrpncgtcqgwqrmqmrwsbmdvqcpggpzcjgtrhvnbpprlwfnjlvrgrdjjvncjmmflrpczqnlbqczggssqfcjcrghqlsztpdjpbbfcdjzzdjbljlsczdmrgshdscvhnltrwchjcmjzmjbhldnqzwwpswsnsldcbdcdpdgpjgwrnfbcpjtvzlhvgggldcfqcwwppvltsvsmwzwdgmhnfggtgtwldmrglcvmstgmbgbjhwwhgdhfrbjlwfhjfppdmffblbbzrplhlhqlsrnsthvjtglqntzcncmvhqnlsvvrscrhlncgncjswfgcvgjlwsndzmsbhbdqsggdgrsfqtzwnjpsdlbsqjrrjwwlbptwqpfqwvpsrtrmstddzdbjjlwvfqcpfczpsnmcgbfpbwcpljdhrgsqftbnplccjphwsdbprqfqqfcvtcdznnhrbdqpccphvdgtspmzzdbnslnrvtfrtbhcfzfgmhrttmdpwftvccjbllhqgtmpgwvbdjgtvtfbfnnsfgzqjmrpbcmqhpfbstznbvgtbhnwbbnjfsthdgdrpfdtvrtmgbwzqqpnlltbshjvnhsmqhqwzgbsfqqccfznjtnnzdrgcvwnlffdgvvqzvwhwfswdmqlrsglntzsnwjzgrhhwzzshwsfmlmrbnmrqlptmjgmtqctrmddzghsgtrbbcsmhtcnrzwmvrjmrnmhbjmflrclvlbzwbgmtnmwqgfmbbnnrdvhqcflglvzbmjzjtvnmrbgghnccfrphjshsgtrhfmmghhpwgclfvzfbdccsfrlfwtsjjnhlndpwcjdtlllhcsvwrwsbppqwhfcvnsnrvthrsbgmgjhpmjdndmdqdgzfvbqfmgfjrnrjchstrjprfwfnjqblhjdgsstvtpcsvmpbhggnwzncpjdhrcllcghhprhwhfgsqpfzphrdlcbrccglsb"

func PrintResult() {
	fmt.Println("--- Day 6 ---")
	result1, result2 := getResult()

	fmt.Println("Result 1: ", result1)
	fmt.Println("Result 2: ", result2)
}

func getResult() (int, int) {
	result := 0
	result2 := 0

	result = parseString(input, 4)
	result2 = parseString(input, 14)

	return result, result2
}

func parseString(str string, marker int) int {
	set := make(map[string]int, marker)
	for i := 0; i < len(str); i++ {
		if len(set) == marker {
			return i
		}
		currRune := string(rune(str[i]))
		if position, ok := set[currRune]; ok {
			set = make(map[string]int, marker)
			i = position + 1
			currRune = string(rune(str[i]))
		}
		set[currRune] = i
	}
	return 0
}
