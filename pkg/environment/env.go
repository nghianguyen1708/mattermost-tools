package environment

type Environment string

const (
	Production Environment = "production"
)

func (e Environment) IsProduction() bool {
	return e == Production
}
