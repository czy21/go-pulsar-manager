import Layout from '@/layout/Layout'
import ClusterList from '@v/cluster/component/ClusterList'
import {renderRoutes, RouteConfig} from "react-router-config";

const PATH = '/cluster'

let routes: RouteConfig[] = [
    {
        path: PATH,
        component: Layout,
        routes: [
            {
                path: PATH,
                component: ClusterList,
                exact: true
            }
        ]
    }
]
export let App = () => renderRoutes(routes)
