package politics

// Ideology represents a political orientation
type Ideology struct {
	name       string
	economy    int8
	diplomacy  int8
	government int8
	society    int8
}

// Ideologies contains the collection of all known ideologies
var Ideologies = []Ideology{
	{
		name:       "Anarcho-Communism",
		economy:    100,
		diplomacy:  50,
		government: 100,
		society:    90,
	},
	{
		name:       "Libertarian Communism",
		economy:    100,
		diplomacy:  70,
		government: 80,
		society:    80,
	},
	{
		name:       "Trotskyism",
		economy:    100,
		diplomacy:  100,
		government: 60,
		society:    80,
	},
	{
		name:       "Marxism",
		economy:    100,
		diplomacy:  70,
		government: 40,
		society:    80,
	},
	{
		name:       "De Leonism",
		economy:    100,
		diplomacy:  30,
		government: 30,
		society:    80,
	},
	{
		name:       "Leninism",
		economy:    100,
		diplomacy:  40,
		government: 20,
		society:    70,
	},
	{
		name:       "Stalinism/Maoism",
		economy:    100,
		diplomacy:  20,
		government: 0,
		society:    60,
	},
	{
		name:       "Religious Communism",
		economy:    100,
		diplomacy:  50,
		government: 30,
		society:    30,
	},
	{
		name:       "State Socialism",
		economy:    80,
		diplomacy:  30,
		government: 30,
		society:    70,
	},
	{
		name:       "Theocratic Socialism",
		economy:    80,
		diplomacy:  50,
		government: 30,
		society:    20,
	},
	{
		name:       "Religious Socialism",
		economy:    80,
		diplomacy:  50,
		government: 70,
		society:    20,
	},
	{
		name:       "Democratic Socialism",
		economy:    80,
		diplomacy:  50,
		government: 50,
		society:    80,
	},
	{
		name:       "Revolutionary Socialism",
		economy:    80,
		diplomacy:  20,
		government: 50,
		society:    70,
	},
	{
		name:       "Libertarian Socialism",
		economy:    80,
		diplomacy:  80,
		government: 80,
		society:    80,
	},
	{
		name:       "Anarcho-Syndicalism",
		economy:    80,
		diplomacy:  50,
		government: 100,
		society:    80,
	},
	{
		name:       "Left-Wing Populism",
		economy:    60,
		diplomacy:  40,
		government: 30,
		society:    70,
	},
	{
		name:       "Theocratic Distributism",
		economy:    60,
		diplomacy:  40,
		government: 30,
		society:    20,
	},
	{
		name:       "Distributism",
		economy:    60,
		diplomacy:  50,
		government: 50,
		society:    20,
	},
	{
		name:       "Social Liberalism",
		economy:    60,
		diplomacy:  60,
		government: 60,
		society:    80,
	},
	{
		name:       "Christian Democracy",
		economy:    60,
		diplomacy:  60,
		government: 50,
		society:    30,
	},
	{
		name:       "Social Democracy",
		economy:    60,
		diplomacy:  70,
		government: 60,
		society:    80,
	},
	{
		name:       "Progressivism",
		economy:    60,
		diplomacy:  80,
		government: 60,
		society:    100,
	},
	{
		name:       "Anarcho-Mutualism",
		economy:    60,
		diplomacy:  50,
		government: 100,
		society:    70,
	},
	{
		name:       "National Totalitarianism",
		economy:    50,
		diplomacy:  20,
		government: 0,
		society:    50,
	},
	{
		name:       "Global Totalitarianism",
		economy:    50,
		diplomacy:  80,
		government: 0,
		society:    50,
	},
	{
		name:       "Technocracy",
		economy:    60,
		diplomacy:  60,
		government: 20,
		society:    70,
	},
	{
		name:       "Centrist",
		economy:    50,
		diplomacy:  50,
		government: 50,
		society:    50,
	},
	{
		name:       "Liberalism",
		economy:    50,
		diplomacy:  60,
		government: 60,
		society:    60,
	},
	{
		name:       "Religious Anarchism",
		economy:    50,
		diplomacy:  50,
		government: 100,
		society:    20,
	},
	{
		name:       "Right-Wing Populism",
		economy:    40,
		diplomacy:  30,
		government: 30,
		society:    30,
	},
	{
		name:       "Moderate Conservatism",
		economy:    40,
		diplomacy:  40,
		government: 50,
		society:    30,
	},
	{
		name:       "Reactionary",
		economy:    40,
		diplomacy:  40,
		government: 40,
		society:    10,
	},
	{
		name:       "Social Libertarianism",
		economy:    60,
		diplomacy:  70,
		government: 80,
		society:    70,
	},
	{
		name:       "Libertarianism",
		economy:    40,
		diplomacy:  60,
		government: 80,
		society:    60,
	},
	{
		name:       "Anarcho-Egoism",
		economy:    40,
		diplomacy:  50,
		government: 100,
		society:    50,
	},
	{
		name:       "Nazism",
		economy:    40,
		diplomacy:  0,
		government: 0,
		society:    5,
	},
	{
		name:       "Autocracy",
		economy:    50,
		diplomacy:  20,
		government: 20,
		society:    50,
	},
	{
		name:       "Fascism",
		economy:    40,
		diplomacy:  20,
		government: 20,
		society:    20,
	},
	{
		name:       "Capitalist Fascism",
		economy:    20,
		diplomacy:  20,
		government: 20,
		society:    20,
	},
	{
		name:       "Conservatism",
		economy:    30,
		diplomacy:  40,
		government: 40,
		society:    20,
	},
	{
		name:       "Neo-Liberalism",
		economy:    30,
		diplomacy:  40,
		government: 40,
		society:    20,
	},
	{
		name:       "Classical Liberalism",
		economy:    30,
		diplomacy:  60,
		government: 60,
		society:    80,
	},
	{
		name:       "Authoritarian Capitalism",
		economy:    20,
		diplomacy:  30,
		government: 20,
		society:    40,
	},
	{
		name:       "State Capitalism",
		economy:    20,
		diplomacy:  50,
		government: 30,
		society:    50,
	},
	{
		name:       "Neo-Conservatism",
		economy:    20,
		diplomacy:  20,
		government: 40,
		society:    20,
	},
	{
		name:       "Fundamentalism",
		economy:    20,
		diplomacy:  30,
		government: 30,
		society:    5,
	},
	{
		name:       "Libertarian Capitalism",
		economy:    20,
		diplomacy:  50,
		government: 80,
		society:    60,
	},
	{
		name:       "Market Anarchism",
		economy:    20,
		diplomacy:  50,
		government: 100,
		society:    50,
	},
	{
		name:       "Objectivism",
		economy:    10,
		diplomacy:  50,
		government: 90,
		society:    40,
	},
	{
		name:       "Ultra-Capitalism",
		economy:    0,
		diplomacy:  40,
		government: 50,
		society:    50,
	},
	{
		name:       "Anarcho-Capitalism",
		economy:    0,
		diplomacy:  50,
		government: 100,
		society:    50,
	},
}