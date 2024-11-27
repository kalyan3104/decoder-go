package main

import (
	"fmt"
	"strings"
)

func main() {
	// Input data
	data := `erd1qqqqqqqqqqqqqpgqyykpxr693gc3mpvvc4k3frq8m7eeguuxqr3q98aec4	00000000000000000500212c130F458a311d858cc56D148C07dFB394738600e2
erd1qqqqqqqqqqqqqpgqz0n6qdgvrz4z23rllustjaqa63cegza40fusz30jca	0000000000000000050013E7A0350c18aA25447fff20b9741Dd471940bb57a79
erd1qqqqqqqqqqqqqpgqz0wkk0j6y4h0mcxfxsg023j4x5sfgrmz0n4s4swp7a	0000000000000000050013dd6B3E5A256eFdE0c93410F546553520940f627ceB
erd1qqqqqqqqqqqqqpgqz0ycyug2rqtpyrh5p33y9vqjv95s3xmaqpnq7uz3qq	0000000000000000050013C982710A1816120eF40C6242B0126169089B7D0066
erd1qqqqqqqqqqqqqpgqz2aa6uaptw8ryf254u93xqwl4lkvu9vmqzpsxshu57	0000000000000000050012BBdD73a15b8e322554AF0b1301dFaFecCe159B0083
erd1qqqqqqqqqqqqqpgqz2ctz77j9we0r99mnhehv24he3pxmsmq0n4sntf4n7	0000000000000000050012B0B17bd22bB2F194Bb9dF3762AB7CC426Dc3607CEB
erd1qqqqqqqqqqqqqpgqz2tgn80j5p5hqh4hx69uc27gz0drcjmmg20skf0wru	000000000000000005001296899df2a069705eB7368Bcc2BC813Da3c4B7b429f
erd1qqqqqqqqqqqqqpgqz35q0ecvzpzhyvum0etghqfcq7d6lsztqqasganmau	00000000000000000500146807E70C104572339B7E568B8138079baFC04B003B
erd1qqqqqqqqqqqqqpgqz5gjr6e8c2qcpqawepr0t2urnkq07p4aqr0q2wwhgd	00000000000000000500151121eb27c2818083Aec846f5ab839D80ff06Bd00dE
erd1qqqqqqqqqqqqqpgqz5jh90tx3d45rjxd8yqkmgt69dqp5ckeyl5s34wdmf	00000000000000000500152572BD668B6B41c8Cd39016Da17A2B401A62d927E9
erd1qqqqqqqqqqqqqpgqz5rth288cy2njyetxeyeqqjv08g4ftn2jpmq2sr3jd	000000000000000005001506bBA8E7C11539132B364990024C79d154ae6A9076
erd1qqqqqqqqqqqqqpgqz7s9ut07qm9uh2rtzzuntnnuxrpejjscqp6q03wqv4	0000000000000000050017A05e2dFe06cBcbA86b10b935ce7C30c3994a180074
erd1qqqqqqqqqqqqqpgqz88aj0jnyk2245qdckqlr0u3esrkm227up4sduprh7	0000000000000000050011CFd93e532594aAD00dc581F1bF91Cc076da95eE06b
erd1qqqqqqqqqqqqqpgqz8q2ggskllzh9dp5988uz42jn95pqeghu00s0mlgca	0000000000000000050011C0a42216FFC572b43429CfC155529968106517E3DF
erd1qqqqqqqqqqqqqpgqzaxt2e4630kpp6uxg5hl6yfefls72us0kklsxhl8fk	00000000000000000500174CB566bA8bEC10EB86452ffd11394Fe1E5720fb5Bf
erd1qqqqqqqqqqqqqpgqze3nszzj6ngfw3g0thm30r3v4tzudexdqz5q6kztm5	000000000000000005001663380852d4d097450F5dF7178E2CaAc5C6e4CD00a8
erd1qqqqqqqqqqqqqpgqzhpcdd8jg77m06zwqmhgw9xdmukn6pfeh2uslry9u8	0000000000000000050015C386b4F247bdb7e84E06eE8714CddF2d3D0539bAb9
erd1qqqqqqqqqqqqqpgqzjctu8xrgn8jmfp503tajjvzz2zq60v92jpsslkh5a	0000000000000000050014b0BE1cc344CF2DA4347c57D9498212840D3D855483
erd1qqqqqqqqqqqqqpgqzjstzefzeastekqpxz6ysh3d0k9eap9aexkscvsdqn	0000000000000000050014A0B16522cF60bCD80130b4485E2d7d8B9E84bDc9aD
erd1qqqqqqqqqqqqqpgqzps75vsk97w9nsx2cenv2r2tyxl4fl402jpsx78m9j	000000000000000005001061eA32162F9C59c0cAc666c50D4b21bf54FeAF5483
erd1qqqqqqqqqqqqqpgqzq2szpgrh95nj6quqgre2qs3x9hct827qqcq28x3p2	000000000000000005001015010503b96939681C0207950211316f859d5E0030
erd1qqqqqqqqqqqqqpgqzqlavpqgwkrnslzrl05aqms2au2327asexksc4u3qg	00000000000000000500103FD604087587387C43fBE9D06E0aef15157bb0c9AD
erd1qqqqqqqqqqqqqpgqzqsrh358v57maeclstufm4w2wrsc52cq0n4s38vfzd	0000000000000000050010203BC687653DbEE71f82f89DD5cA70e18a2b007CeB
erd1qqqqqqqqqqqqqpgqzqvm5ywqqf524efwrhr039tjs29w0qltkklsa05pk7	000000000000000005001019BA11c00268aaE52e1DC6f89572828AE783EbB5bF
erd1qqqqqqqqqqqqqpgqzs3y0vr4t4sqyet32t6jp9m85l3hlr93u00sdj9rxa	00000000000000000500142247b0755D6002657152f5209767A7E37f8Cb1E3Df
erd1qqqqqqqqqqqqqpgqzsqek20mul9t07kwj8fswjl85qnmysweqrvs4u6qq3	0000000000000000050014019B29fbE7CaB7faCE91d3074be7A027b241d900D9
erd1qqqqqqqqqqqqqpgqzt7clww7329ajkq07x295h7hvl2hgnsaexks5m7sen	0000000000000000050012fD8fb9de8a8bd9580Ff1945A5fd767d5744e1DC9aD
erd1qqqqqqqqqqqqqpgqzud07wsp6ghwsr8rwdhwt3pgrq6aveelznyq3h0yft	00000000000000000500171afF3A01D22EE80ce3736ee5c4281835d6673f14C8
erd1qqqqqqqqqqqqqpgqzuxy0ve8scj6jyl82gzy9m26lmd6m69kg20s5gt357	00000000000000000500170c47B3278625A913E7520442ED5aFedBADe8b6429F
erd1qqqqqqqqqqqqqpgqzvcvh8ur47nlwxsa9wf7gvrlmf764aerczfsnsjudc	000000000000000005001330cb9f83Afa7f71A1d2B93E4307fdA7daAf723C093
erd1qqqqqqqqqqqqqpgqzw0c6t3puxg4rfw9ae8zc2xl3wvmqwxmqprsucygxr	00000000000000000500139f8d2E21E19151a5C5Ee4e2c28df8B99b038db0047
erd1qqqqqqqqqqqqqpgqzw0d0tj25qme9e4ukverjjjqle6xamay0n4s5r0v9g	00000000000000000500139eD7ae4aa03792E6bcb332394A40fE746EEfa47CeB
erd1qqqqqqqqqqqqqpgqzwl6axaq2kfjaglfsrm3eqc0rs3cn85rqrast276y9	0000000000000000050013BFaE9bA055932EA3E980f71C830F1C23899e8300Fb
erd1qqqqqqqqqqqqqpgqzwv3sn2h4juyya2aadacmrt4hqj9r7a70n4shpwswh	000000000000000005001399184d57aCB842755deb7B8D8d75b82451FBBE7ceB
erd1qqqqqqqqqqqqqpgqzy0qwv0jf0yk8tysll0y6ea0fhpfqlp32jpsder9ne	00000000000000000500111E0731f24BC963Ac90FfDe4D67Af4dc2907C315483`

	// Split input data into lines
	lines := strings.Split(data, "\n")

	// Iterate over each line and process
	for _, line := range lines {
		parts := strings.Split(line, "\t")
		if len(parts) == 2 {
			// Extract and print the binary data
			fmt.Println(parts[1])
		}
	}
}