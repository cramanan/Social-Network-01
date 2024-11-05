import React from "react";
import Users from "./Users";
import SearchBar from "./SearchBar";

const FindUser = () => {
    return (
        <>
            <div className="w-72 h-96 flex flex-col bg-white bg-opacity-40 rounded-3xl gap-2 p-3">
                <input
                    id="to-input"
                    type="text"
                    className="outline-none bg-white/0 placeholder:text-black"
                    placeholder="To"
                />

                <SearchBar id="seach-bar-user" />

                <span>Suggestions</span>

                <div className="flex flex-col gap-2 mx-5">
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                    <Users />
                </div>
            </div>
        </>
    );
};

export default FindUser;
