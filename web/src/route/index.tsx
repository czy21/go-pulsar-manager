import {RouteConfig} from "react-router-config";
import {App as CLUSTER} from '@v/cluster'

const routes: RouteConfig[] = [
    {
        path: "/cluster",
        component: CLUSTER,
    }
];
export default routes