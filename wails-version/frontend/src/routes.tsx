import { Route, Routes, HashRouter } from "react-router-dom";

import PageListTransactions from "./listTransactions";
import NewTransaction from "./newTransaction";

const WalletRoutes = () => {
    return (
        <HashRouter basename={"/"}>
            <Routes>
                <Route Component={PageListTransactions} path="/" />
                <Route Component={NewTransaction} path="/newtransaction" />
            </Routes>
        </HashRouter>
    )
}

export default WalletRoutes;