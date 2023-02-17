package prompts

import (
	"errors"
	"fmt"

	"github.com/pterm/pterm"
)

// Errors
var (
	NoServicesDefined = errors.New("no services defined in project config")
	NoValidDomains    = errors.New("no valid domains")
)

const (
	// Prompts
	BranchPrompt        = "Branch:"
	CommitMessagePrompt = "Commit Message:"
	DescriptionPrompt   = "Description:"
	EntryPointPrompt    = "Entry Point:"

	MemoryPrompt           = "Memory:"
	PathsPrompt            = "Paths [comma separated]:"
	SizePrompt             = "Size:"
	TTLPrompt              = "TTL:"
	UnitTypePrompt         = "Unit type:"
	TrueSelect             = "True"
	TrueSelectL            = "true"
	FalseSelect            = "False"
	FalseSelectL           = "false"
	TagsPrompt             = "Tags [comma separated]:"
	RegexPrompt            = "Use Regex For Match:"
	LocalPrompt            = "Local:"
	PathPrompt             = "Path:"
	GenerateRepoPrompt     = "Generate a Repository:"
	UseTemplatePrompt      = "Use a Template:"
	SelectTemplatePrompt   = "Select a Template:"
	RepositoryNamePrompt   = "Repository Name:"
	RepositorySelectPrompt = "Select a Repository"
	ClonePrompt            = "Clone this Repository"
	Domains                = "Domains:"
	EmbedTokenPrompt       = "Embed Git Token Into Clone URL:"
	BranchSelectPrompt     = "Select a Branch:"
	PrivatePrompt          = "Private:"
	SourcePrompt           = "Code source:"
	CodeLanguagePrompt     = "Code Language:"
	CallPrompt             = "Entry Point:"
	TimeoutPrompt          = "Time To Live:"

	NoDomainGeneratePrompt = "No domains found, generate one?"

	// Error messages
	Required                   = "Required"
	MustBeABooleanValue        = "Must be a boolean value"
	FieldNotDefinedInConfig    = "field not defined in config: %#v"
	DoubleStringNotFound       = "%s %s not found"
	StringIsRequired           = "%s is required"
	StringIsNotAValidSelection = "`%s` not a valid selection: %v"
	InvalidSize                = "Invalid size: %s Ex:(10, 10GB, 10PB)"
	SelectPromptNoOptions      = "no options to select from for prompt: %s"
	NoServiceFromFlag          = "unable to find service with selection: (--%s %s)"

	//    Device Errors
	TagLessThanThreeCharacters = "Tags cannot be less than three characters"
)

var (
	SelectionNone = "(none)"
)

func PanicIfPromptNotEnabled(prompt string) {
	panicIfPromptNotEnabled(prompt)
}

func panicIfPromptNotEnabled(prompt string) {
	if PromptEnabled == false {
		pterm.Warning.Printfln("Failed to prompt: %s", prompt)
		panic("Prompting when prompt not enabled")
	}
}

func panicIfPromptNotEnabledSelection(val string, prompt string, opts []string) {
	if PromptEnabled == false {
		pterm.Warning.Printfln("%s not a valid selection %v", val, opts)
		panic(fmt.Sprintf("Prompt with prompt disabled, %s", prompt))
	}
}

func panicIfPromptNotEnabledError(err error) {
	if PromptEnabled == false {
		panic(fmt.Sprintf("Prompt with prompt disabled: %s", err))
	}
}