package serpent

//@todo Adde SerpentKeyGenertor later
type KeyGeneratorParameters struct {
	alpha   uint8
	numbers uint8
	special uint8
}

type SerpentStringKeyGenerator interface {
	GenerateKey(generatorParams *KeyGeneratorParameters) string
}

func (serpent *Serpent) GenerateKey(generatorParams *KeyGeneratorParameters) string {
	return ""
}
