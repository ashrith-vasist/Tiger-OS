package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"tiger/pkg/basic"
)

var calc = &cobra.Command {
	Use: "Calculator Command",
	Short: "The Calculator Command runs a simple calculator command, the first version of bc.",
	Version: version,
	Run: func(cmd *cobra.Command, args[] string) {
		operand1, _ := cmd.Flags().GetString("operand1")
		operand2, _ := cmd.Flags().GetString("operand2")
		operator, _ := cmd.Flags().GetString("operator")

		basic.Calculate(operand1,operand2,operator)
	},
}

func init() {
	rootCmd.AddComand(calc)
	calc.Flags().StringP("operand1","o1",viper.GetString("operand1"),"Get Operand 1")
	calc.Flags().StringP("operand2","o2",viper.GetString("operand2"),"Get Operand 2")
	calc.Flags().StringP("operator","op",viper.GetString("operator"),"Get Operator")

	calc.MarkFlagRequired("operand1")
	calc.MarkFlagRequired("operand2")
	calc.MarkFlagRequired("operator")
}