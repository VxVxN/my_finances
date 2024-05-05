export enum OrderType {
    BUY = 'buy',
    SELL = 'sell'
}

export interface Order {
    id:             number;
    type:           OrderType;
    datetime:       Date;
    baseCurrency:  string;
    value:          number;
    quoteCurrency: string;
    price:          number;
}

