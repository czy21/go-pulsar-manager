import React from "react";
import react from "react";
import stub from "@/init";
import styles from './index.less'
import {MenuFoldOutlined, MenuUnfoldOutlined} from '@ant-design/icons';
import locale from "@/locale";

const Header: React.FC<any> = (props: any) => {
    const homeState = stub.ref.reactRedux.useSelector((state: any) => state.home)
    react.useEffect(() => {
        stub.api.post('option/query', {keys: ["environment"]})
            .then((t: any) => {
                const e = t.data.data["environment"]
                stub.store.dispatch(stub.reducer.action.home.setEnvironments(e))
                if (e) {
                    stub.store.dispatch(stub.reducer.action.home.setEnvironment(e[0]))
                }
            })
    }, [])
    return (
        <stub.ref.antd.Layout.Header className={styles.header}>
            <stub.ref.antd.Row justify="space-between">
                <stub.ref.antd.Col>
                    {React.createElement(homeState.collapsed ? MenuUnfoldOutlined : MenuFoldOutlined, {
                        className: styles.collapse,
                        onClick: () => {
                            stub.store.dispatch(stub.reducer.action.home.collapse())
                        },
                    })}
                </stub.ref.antd.Col>
                <stub.ref.antd.Col>
                    <stub.ref.antd.Select
                        options={homeState.environments}
                        defaultValue={homeState.environment.value}
                        key={homeState.environment.value}
                    />
                    <stub.ref.antd.Select
                        options={Object.entries<any>(locale || []).map(([k, v]) => {
                            return {
                                label: v["global.language.name"], value: k
                            }
                        })}
                        defaultValue={homeState.locale.key}
                        onSelect={(value: any, option: any) => {
                            stub.store.dispatch(stub.reducer.action.home.switchLocale({key: value, message: {...locale["en_US"], ...locale[`${value}`]}}))
                        }}
                        style={{margin: "0 24px"}}
                    />
                </stub.ref.antd.Col>
            </stub.ref.antd.Row>

        </stub.ref.antd.Layout.Header>
    )

}

export default Header