import Layout from '@/shared/layouts/Layout';
import {OrderList} from '@/widgets/OrderList';
import {PageTitle} from '@/features/PageTitle';

export default function OrdersPage() {
    return (
        <Layout>
            <PageTitle title={'Сделки'}/>
            <OrderList/>
        </Layout>
    );
}