![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

Based on [EFF](https://eff.org/dice) updates to diceware.
PoC to implement the generation method found in the PDF.

1 Step 1: Roll five dice all at once. Note the faces that come up.
2 Step 2: Your results might look like this, reading left to right: 4, 3, 4, 6, 3. Write
those numbers.
3 Step 3: Open EFF's Wordlist (link above) to find the corresponding word next to
43463.
4 Step 4: You will find the word "panoramic." This is the first word in your
passphrase, so write it down.
5 Step 5: Repeat the above steps five more times to come up with a total of SIX words.
When you are done, your passphrase may look something like this: panoramic nectar
precut smith banana handclap.
6 Step 6: Come up with your own mnemonic to remember your phrase. It might be a
story, scenario, or sentence that you will remember and can remind you of the words
you chose, in order.

Wordlist notes:

- Long wordlist should be used with five dice.
- Both short lists should be used with four dice.

Usage: go run dice.go -wordlist=(1-3)

Thank you to [Joseph Bonneau](https://jbonneau.com) and [Arnold G. Reinhold](http://diceware.com/) for the creations.
