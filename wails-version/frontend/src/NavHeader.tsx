import { useEffect } from "react"

function NavHeader() {
    const headerStyle: React.CSSProperties = { display: "flex", alignItems: "center", justifyContent: "space-between", flexWrap: 'wrap' }
    const navStyle: React.CSSProperties = { display: "flex", gap: "10px", flexWrap: 'wrap' }
    
    useEffect(()=>{
        
    },[0])
    return (
        <header style={headerStyle}>
            <h1><a href="/">Wallet</a></h1>
            <nav style={navStyle}>
                <a href="/#/newtransaction">New Transaction</a> |
                <a href="/">List Transactions</a>
            </nav>
        </header >
    )
}

export default NavHeader