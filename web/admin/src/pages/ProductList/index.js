import React, {useState, useEffect} from "react";
import {Avatar, Button, Layout, List, PageHeader, Pagination, Skeleton} from "antd";
import {listProduct} from "../../api/catalog";
import {Link, useRouteMatch} from "react-router-dom";

export default function ProductList() {
    const [productList, setProductList] = useState([]);
    let {path, url} = useRouteMatch();

    useEffect(() => {
        listProduct().then((res) => {
            setProductList(res.data.results)
        });

    }, []);

    return (
        <Layout.Content>
            <PageHeader
                ghost={false}
                onBack={() => window.history.back()}
                title="Products"
            />
            <div style={{background: "#fff"}}>
                <List
                    pagination={<Pagination defaultCurrent={1} total={50}/>}
                    dataSource={productList}
                    renderItem={(item, i) => (
                        <List.Item
                            actions={[<a key={i}>edit</a>]}
                        >
                            <Skeleton avatar title={false} loading={item.loading} active>
                                <List.Item.Meta
                                    avatar={<Avatar shape="square" size={64} src={item.image[0].url}/>}
                                    title={<Link to={`${path}/${item.id}`}>{item.name}</Link>}
                                    description={item.price}
                                />
                            </Skeleton>
                        </List.Item>
                    )}
                />
            </div>
        </Layout.Content>
    )
}
