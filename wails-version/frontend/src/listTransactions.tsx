import { useState } from "react"
import { ListTransactions, DeleteTransaction } from "../wailsjs/go/main/App";

type transactionType = {
    ID: string;
    Title: string;
    Amount: number;
    Categories: string[]
};

function PageListTransactions() {
    const [transactions, setTransactions] = useState<transactionType[]>([]);

    const listTransactions = () => {
        ListTransactions().then()
        ListTransactions().then(ts => { setTransactions(ts == null ? [] : ts) })
    }

    listTransactions()

    return (
        <div>
            <h2>Transactions list</h2>
            <table>
                <thead>
                    <tr>
                        <th>Title</th>
                        <th>Amount</th>
                        <th>Categories</th>
                        <th></th>
                    </tr>
                </thead>
                <tbody>
                    {transactions.map(trans => {
                        return (
                            <tr key={trans.ID}>
                                <td>{trans.Title}</td>
                                <td>{trans.Amount}</td>
                                <td>{trans.Categories.join(", ")}</td>
                                <td><button onClick={() => { DeleteTransaction(trans); listTransactions() }} >DELETE</button></td>
                            </tr>
                        )
                    })}
                </tbody>
            </table>
        </div>
    )
}

export default PageListTransactions