import React from "react"
import stub from '@/init'

export interface MenuModel {
    name: string,
    path?: string
    icon?: React.ReactNode
    children?: Array<MenuModel>
}

const menus: MenuModel[] = [
    {
        name: "集群",
        path: "/cluster",
        icon: <stub.ref.icon.ai.ClusterOutlined/>,
    },
    {
        name: "数据库",
        path: "/db",
        icon: <stub.ref.icon.ai.ContactsOutlined/>,
    },
];
export default menus