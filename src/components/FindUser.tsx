import React from 'react'
import { Search } from "./icons/Search"
import Users from "./Users"

const FindUser = () => {
    return (
        <>
            <div className="w-72 h-96 flex flex-col bg-white bg-opacity-40 rounded-3xl gap-2 p-3">
                <input id="to-input" type="text" className="outline-none bg-white/0 placeholder:text-black" placeholder="To" />
                <div className="flex">
                    <div className="flex justify-center items-center w-12 h-8 border border-white rounded-l-3xl bg-white bg-opacity-40 border-r-0" aria-hidden="true">
                        <Search />
                    </div>
                    <label htmlFor="search-input" className="sr-only">Search</label>
                    <input id="search-user-input" type="search" className="w-80 h-8 border rounded-r-3xl border-l-0  border-white bg-white bg-opacity-40 focus:outline-none" placeholder="Search" aria-label="Search" />
                </div>
                <span >Suggestions</span>
                <div className="flex flex-col gap-2 mx-5">
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                </div>
            </div>
        </>
    )
}

export default FindUser