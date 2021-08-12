##Structure:

#Current State:
    - Orderbook
    - Up/Down volume (sequence of trades on bid or ask?)
    - Sequence of prices

#Model:
    Training Features:
        - Orderbook
        - Up/Down volume (sequence of trades on bid or ask?)
        - Sequence of prices
        - Market Profile?

    Outputs:
        - A distribution of possible next prices could be useful?

#Strategies:
    Market Making:
        - Avelleneda + Stoikov inventory risk management
        - Constantly post best bid/ask
        - Try to avoid big moves while carrying inventory (or bet on the big moves)
            - Use the model to try to predict this
    Trend following:
        - Take +EV bets using the model
    