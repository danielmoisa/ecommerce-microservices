import React, {useState, useEffect} from "react";
import {Avatar, Button, List, PageHeader, Pagination, Skeleton} from "antd";
import {listProduct} from "../../api/catalog";
import {Link, useRouteMatch} from "react-router-dom";

export default function OrderList() {
    const [orderList, setOrderList] = useState([]);
    let {path, url} = useRouteMatch();

    useEffect(() => {
        listProduct().then((res) => {
            setOrderList(res.data.results)
        });

    }, []);

    return (
        <div>
            <PageHeader
                ghost={false}
                onBack={() => window.history.back()}
                title="Order"
            />
            <div style={{background: "#fff"}}>
                <List
                    pagination={<Pagination defaultCurrent={1} total={50}/>}
                    dataSource={orderList}
                    renderItem={(item, i) => (
                        <List.Item
                            actions={[<a key={i}>edit</a>]}
                        >
                            <Skeleton avatar title={false} loading={item.loading} active>
                                <List.Item.Meta
                                    avatar={<Avatar shape="square" size={64} src={item.images[0]}/>}
                                    title={<Link to={`${path}/${item.id}`}>{item.name}</Link>}
                                    description={item.price}
                                />
                            </Skeleton>
                        </List.Item>
                    )}
                />
            </div>
        </div>
    )
}
