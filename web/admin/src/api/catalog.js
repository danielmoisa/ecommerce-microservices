import service from './index'

export const listProduct = (pageNum, pageSize) => {
    return service.get("/admin/v1/catalog/products", {
        pageNum, pageSize
    })
};

export const getProductDetail = (id) => {
    return service.get("/admin/v1/catalog/products" + id)
};
