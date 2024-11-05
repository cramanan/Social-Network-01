import React from "react";
import { SearchIcon } from "./icons/SearchIcon";

const SearchBar = ({ id }: { id: string }) => {
    return (
        <div className="flex">
            <div
                className="flex justify-center items-center w-12 h-8 border border-white rounded-l-3xl bg-white bg-opacity-40 border-r-0"
                aria-hidden="true"
            >
                <SearchIcon />
            </div>
            <label htmlFor={id} className="sr-only">
                Search
            </label>
            <input
                id={id}
                type="search"
                className="w-52 h-8 border rounded-r-3xl pr-2 border-l-0 border-white bg-white bg-opacity-40 focus:outline-none xl:w-80"
                placeholder="Search"
                aria-label="Search"
            />
        </div>
    );
};

export default SearchBar;
