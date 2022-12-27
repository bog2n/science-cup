package main

import (
	"fmt"
	"motorola/ribosome"
)

var genome string = "attaaaggtt tataccttcc caggtaacaa accaaccaac tttcgatctc ttgtagatct gttctctaaa cgaactttaa aatctgtgtg gctgtcactc ggctgcatgc ttagtgcact cacgcagtat aattaataac taattactgt cgttgacagg acacgagtaa ctcgtctatc ttctgcaggc tgcttacggt ttcgtccgtg ttgcagccga tcatcagcac atctaggttt cgtccgggtg tgaccgaaag gtaagatgga gagccttgtc cctggtttca acgagaaaac acacgtccaa ctcagtttgc ctgttttaca ggttcgcgac gtgctcgtac gtggctttgg agactccgtg gaggaggtct tatcagaggc acgtcaacat cttaaagatg gcacttgtgg cttagtagaa gttgaaaaag gcgttttgcc tcaacttgaa cagccctatg tgttcatcaa acgttcggat gctcgaactg cacctcatgg tcatgttatg gttgagctgg tagcagaact cgaaggcatt cagtacggtc gtagtggtga gacacttggt gtccttgtcc ctcatgtggg cgaaatacca gtggcttacc gcaaggttct tcttcgtaag aacggtaata aaggagctgg tggccatagt tacggcgccg atctaaagtc atttgactta ggcgacgagc ttggcactga tccttatgaa gattttcaag aaaactggaa cactaaacat agcagtggtg ttacccgtga actcatgcgt gagcttaacg gaggggcata cactcgctat gtcgataaca acttctgtgg ccctgatggc taccctcttg agtgcattaa agaccttcta gcacgtgctg gtaaagcttc atgcactttg tccgaacaac tggactttat tgacactaag aggggtgtat actgctgccg tgaacatgag catgaaattg cttggtacac ggaacgttct gaaaagagct atgaattgca gacacctttt gaaattaaat tggcaaagaa atttgacacc ttcaatgggg aatgtccaaa ttttgtattt cccttaaatt ccataatcaa gactattcaa ccaagggttg aaaagaaaaa gcttgatggc tttatgggta gaattcgatc tgtctatcca gttgcgtcac caaatgaatg caaccaaatg tgcctttcaa ctctcatgaa gtgtgatcat tgtggtgaaa cttcatggca gacgggcgat tttgttaaag ccacttgcga attttgtggc actgagaatt tgactaaaga aggtgccact acttgtggtt acttacccca aaatgctgtt gttaaaattt attgtccagc atgtcacaat tcagaagtag gacctgagca tagtcttgcc gaataccata atgaatctgg cttgaaaacc attcttcgta agggtggtcg cactattgcc tttggaggct gtgtgttctc ttatgttggt tgccataaca agtgtgccta ttgggttcca cgtgctagcg ctaacatagg ttgtaaccat acaggtgttg ttggagaagg ttccgaaggt cttaatgaca accttcttga aatactccaa aaagagaaag tcaacatcaa tattgttggt gactttaaac ttaatgaaga gatcgccatt attttggcat ctttttctgc ttccacaagt gcttttgtgg aaactgtgaa aggtttggat tataaagcat tcaaacaaat tgttgaatcc tgtggtaatt ttaaagttac aaaaggaaaa gctaaaaaag gtgcctggaa tattggtgaa cagaaatcaa tactgagtcc tctttatgca tttgcatcag aggctgctcg tgttgtacga tcaattttct cccgcactct tgaaactgct caaaattctg tgcgtgtttt acagaaggcc gctataacaa tactagatgg aatttcacag tattcactga gactcattga tgctatgatg ttcacatctg atttggctac taacaatcta gttgtaatgg cctacattac aggtggtgtt gttcagttga cttcgcagtg gctaactaac atctttggca ctgtttatga aaaactcaaa cccgtccttg attggcttga"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Printf("============= Offset %d =============\n", i)
		prot, err := ribosome.GetAminoAcids(genome[i:])
		if err != nil {
			fmt.Print(err)
			return
		}
		for i := range prot {
			fmt.Println(prot[i])
		}
	}
}
