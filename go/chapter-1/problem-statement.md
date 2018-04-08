Create a simple tip calculator. The program should prompt
for a bill amount and a tip rate. The program must compute
the tip and then display both the tip and the total amount
of the bill.

Example output:

What is the bill? $200
What is the tip percentage? 15
The tip is $30.00
The toal is $230.00

Pseudocode Algorithm

TipCalculator
  Initialize billAmount to 0
  Initialize tip to 0
  Initialize tipRate tip to 0
  Initialize total tip to 0

  Prompt for billAmount with "What is the bill amount?"
  Prompt for tipRate with "What is the tip rate?"

  convert billAmount to a number
  convert tipRate to a number

  tip = billAmount * (tipRate / 100)
  round tip up to nearest cent
  total = billAmount + tip

  Display "Tip: $" + tip
  Display "Total: $" + billAmount
End
