import Link from 'next/link'
import React from "react";

const linkStyle = {
    marginRight: 40
}

export default function Header() {
    return (
        <div>
            <Link href="/">
                <a style={linkStyle}>Home</a>
            </Link>
            <Link href="/about">
                <a style={linkStyle}>About</a>
            </Link>
        </div>
    )
}