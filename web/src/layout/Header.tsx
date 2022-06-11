import React from "react";
import react from "react";
import stub from "@/init";
import styles from './index.less'
import {MenuFoldOutlined, MenuUnfoldOutlined} from '@ant-design/icons';
import locale from "@/locale";

export const mapStateToProps = (state: any) => {
    return {
        collapsed: state.home.collapsed,
        locale: state.home.locale
    }
};
const Header: React.FC<any> = (props: any) => {
    react.useEffect(() => {
        console.log(props.locale)
    }, [])
    return (
        <stub.ref.antd.Layout.Header className={styles.header}>
            <stub.ref.antd.Row justify="space-between">
                <stub.ref.antd.Col>
                    {React.createElement(props.collapsed ? MenuUnfoldOutlined : MenuFoldOutlined, {
                        className: styles.collapse,
                        onClick: () => {
                            stub.store.dispatch(stub.reducer.action.home.collapse())
                        },
                    })}
                </stub.ref.antd.Col>
                <stub.ref.antd.Col>
                    <stub.ref.antd.Select
                        options={Object.entries(locale || []).map(([k, v]) => {
                            return {
                                label: (v as any)["global.language.name"], value: k
                            }
                        })}
                        defaultValue={props.locale.key}
                        onSelect={(value: any, option: any) => {
                            let message = {...locale["en_US"]}
                            Object.entries(locale[`${value}`])
                                .forEach(([k, v]) => {
                                    message[k] = v
                                })
                            stub.store.dispatch(stub.reducer.action.home.switchLocale({key: value, message: message}))
                        }}
                        style={{margin: "0 24px"}}
                    />
                </stub.ref.antd.Col>
            </stub.ref.antd.Row>

        </stub.ref.antd.Layout.Header>
    )

}

export default stub.ref.reactRedux.connect(mapStateToProps)(Header)