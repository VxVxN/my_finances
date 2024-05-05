import React, {FC} from 'react';
import {Card, CardBody, CardHeader} from '@nextui-org/card';
import './styles.scss';
import {Order, OrderType} from '@/entities/OrderCard';
import {DateUtils} from '@/shared/utils/DateUtils';
import {Chip} from '@nextui-org/chip';

interface OrderCardProps {
    data: Order;
}

export const OrderCard: FC<OrderCardProps> = ({data}) => {
    const TypeOrder = (orderType: OrderType) => {
       return (<Chip variant="flat" size="sm" color={orderType === OrderType.BUY ? 'danger' : 'success'}>{orderType.toUpperCase()}</Chip>)
    };

    console.log(data);

    return (
        <Card className="order-card" isPressable >
            <CardHeader className='head'>
                <div className={`title ${data.type}`}>
                    {data.baseCurrency}/{data.quoteCurrency}
                </div>
                <div className='date'>
                    {DateUtils.formatDate(data.datetime)}
                </div>
            </CardHeader>
            <CardBody className='body'>
                <div className="amount">
                    <div className='price'>
                        {data.price}
                    </div>
                    <div style={{fontSize: '.75rem'}}>x</div>
                    <div className="value">{data.value}</div>
                </div>
                {TypeOrder(data.type)}
            </CardBody>
        </Card>
    );
};
