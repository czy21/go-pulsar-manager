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
    }
];
export default menus