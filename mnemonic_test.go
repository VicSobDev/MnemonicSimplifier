package mnemonic_simplifier_test

import (
	"testing"

	mnemonic_simplifier "github.com/vicsobdev/MnemonicSimplifier"
)

var result = []mnemonic_simplifier.Account{
	{Address: "0xDFbc2DFA78a113fa03dddf7c030362900B0fa206", PublicKey: "0x02e1662ae2c582bc0ff6ed6da5094e9d80fd33c8aa111ec1388acaa5fee1ed9989", PrivateKey: "0x1b2cde86de16f9ec7611c4d927b7a6e71d4e20817155537e3aeb081c55e3a54e"},
	{Address: "0xD4F27D86433Ae4991370dabDCacCC52a8937032E", PublicKey: "0x0336ed3cf916b019d44996abbe4b912cd3ab6837525919e152e8cb90873d18d3d1", PrivateKey: "0x8a7e36cfd2d597041c398868eeb26b1b56264cb1deea868c7b8dc59c78d05796"},
	{Address: "0x4cEc80747e5c717144748a0f1eacd81A1Dc1c085", PublicKey: "0x035cf32127c3716b9b27ada52e6e1e6275f76e511677a300ef4cfe7c1c0ff93d18", PrivateKey: "0xd5271aa81b94d28fa89f31a238c4c2e46025be371a26dbcca6d16b01b92dee83"},
	{Address: "0xcE4f6141358f6b0d0caa882C39a3B927cc9FA403", PublicKey: "0x034502d4b7be871e57c41343a60a497c30f09d99d560bfb0269190c2e0ddd364ca", PrivateKey: "0x6dbb2930ed16727faefad0e57e8d34d8c3d1bf67d983c3648f30b9655177dc27"},
	{Address: "0x674DA4245aDfEF605A27506f3cD5046EB7E8e5fB", PublicKey: "0x0398de5a1cc79427bc1bf3e77280d48a50f8d74e2ae50047ab1e06ba91e3f988f2", PrivateKey: "0xd832d5f0d6737dc2d5c54a7491d5496b4a7121f6c2e5bafd3361e51c9a28fede"},
	{Address: "0xA713767Dfba82f3E3423Ec8bb3991476ED7a503c", PublicKey: "0x0355167a98f28b98dacd8f849ac6432ad15cec3f8212a78e886a4cba4350894679", PrivateKey: "0x5dad860f93cc9cd0467545e66b1928ec8275f2cd6bc98156c3075de0eb194c26"},
	{Address: "0xf23DB682B027DDc2495eaBA308b829aD3771E722", PublicKey: "0x0393db81d3a40697cc0ab2292c5fde52877db58ca80e242f0a8414dd804685dc0b", PrivateKey: "0xe8cc2ce5bc6c3124b2b147f11c08f9ff4e0ffef6de8a72935ac17de1447e5b91"},
	{Address: "0xA389DC88cE611b6c0FB1902138347791a0Dd1215", PublicKey: "0x02b876f7518da6a0edad501b58794add3f017284da25a404bd2fb14ed80d45a361", PrivateKey: "0xd8603336fa6c66cd61f9cc1fa26350552eb92ea3484a202596579af1fe7ed213"},
	{Address: "0x5004a24296f5ab565aaa676598B37311d62E9F05", PublicKey: "0x02fc690b439d446266b21cb4ae55772e2940e2e3c3f76d89baba784c0c06186bed", PrivateKey: "0x00fd3604a4e1f5ef9d61420a4000d221f2463b73f6b948f2a2000529134fd7a3"},
	{Address: "0x003fEC256c5f293ce63eE3DF40fa82a08d945Ce5", PublicKey: "0x036271172731f20ed524a40e3136025c5874302ec45e7b46b3aef5bf68eb880a8c", PrivateKey: "0x728bff3bf34134c8bd395dd2e07a1e20713bf15f0b9302ebcb26390e79b86ea5"},
	{Address: "0xdCb6205bC8C1434FbBd94EDd5c6C4Be5E4A7b577", PublicKey: "0x03c76a38e8fb32e99a1f3f80088195b7dab90eb4b282262a693ab5f5ab6bbc2b9d", PrivateKey: "0xc9ae58e4f612dc46a8dbb15158626af6ff56b40d0e90a69d045fd03f72c088ed"},
	{Address: "0xEDecbcb02Cadbb26138196F3E24769C547C4c639", PublicKey: "0x027d0b6e7bf2596659f9f64dd49ad43e46f58202245792dfd24e0eb01907f7c4df", PrivateKey: "0xb2b25a0bc41ebf6eee5f34f05ad659814ee0f7d6dbceff6a2cb5bace8e08e0a3"},
	{Address: "0x5062b3097e497B7087ebB9391Ef7747bAB7B899a", PublicKey: "0x0305dea86b59231c26c32688ad8dd0e24707bc4521947c5b2ac3fb51cf26dde35c", PrivateKey: "0x351f60ff6efe0070f1f66f746d2cd58848276b9623c9fca5203956ddced8a92a"},
	{Address: "0x9ad96264fFCcDfFE6daB7f898C911896B2c846AE", PublicKey: "0x023bc83ffd0226cb6653ff2d5c5ee971295d54f7bf955c1e9db7383d976d8a834c", PrivateKey: "0x8156fff0f5f3d1f0be67760c059c37e33f3d433bcba49d541b23f0956523c32d"},
	{Address: "0xE65eae7B51566327D12A99d6C86513b2a7420f90", PublicKey: "0x02ef92e2383ecc6b015a5e11717b88717fb63b3aeed8e52aef09b80808cccb04f0", PrivateKey: "0xbcde8b6a047e26d7d8f1937685375123b6af5c2a1e13f7989c53679c04a55a7b"},
	{Address: "0x6aF37906A6b4d98b702A4cDF563dD61a1535853f", PublicKey: "0x032a7ce81f6233683869a2af011d6586115d87c23a6ec081185078d86506b013b1", PrivateKey: "0xd6e74034c347c6967337dd87ad85b39ceb5dfd72d22ab3dc199fa61183cfbe1b"},
	{Address: "0xb2e7B3f282f442616cE62d8E53f418342684266C", PublicKey: "0x03ac0a8103d0d1c22da406fa48b9f23fd4292ba3010b6ab1becc7120bc98fa639e", PrivateKey: "0xbd88fa5c975be4b0ed475ff2de76acd9379c5594c1d75a730c3cad621cf876d0"},
	{Address: "0xED5AC3a920Ce6A35E14d8A91F13874308Db0875E", PublicKey: "0x03aa1ac0ce3b77da76c671b56c218e595309468992d937e68357915ec9db3f9d82", PrivateKey: "0x0631029755f64be8da01ce2627d8c98fdc15398569da5b5c37fc07b9b3bab32b"},
	{Address: "0x4f4D9509180805C5eb13171AC089B5b3E9C98642", PublicKey: "0x0378a380b8b016cea040439dbb99eb05d39dc9d5c833c03b42ab33cfa6a9ba5d11", PrivateKey: "0x3e643b11c8f845a86f1cdcb1e7185ded7ef025c70c013f6c3ec538e9b34706c6"},
	{Address: "0x0E4f5E7F1f953F0f3696A92cca19E0Ccc1a38322", PublicKey: "0x03522cb62e0b042245a43568b3334aaad3b19486ad4bdc2de9445848a9f1694f8b", PrivateKey: "0x50fc7e4a5291c130d5525b8391793d9dc31884e2f8c8dd45dd5ba1889c3f6dc1"},
}

func TestNewWallet(t *testing.T) {
	mnemonic := "admit initial slice detail glue they chair office misery wish pave fitness"
	wallet, err := mnemonic_simplifier.NewWallet(mnemonic)
	if err != nil {
		t.Errorf("Error creating wallet: %s", err)
	}

	if len(wallet.Accounts) != 10 {
		t.Errorf("Expected 10 accounts, got %d", len(wallet.Accounts))
	}

	for i, account := range wallet.Accounts {
		if account.Address != result[i].Address {
			t.Errorf("Expected address %s, got %s", result[i].Address, account.Address)
		}
	}

}
