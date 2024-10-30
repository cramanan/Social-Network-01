import React from 'react'
import { HomeIcon } from "./icons/HomeIcon"

const MobileBottomNav = () => {
    return (
        <>
            <div className="relative w-full h-full z-60">
                <nav className="relative flex flex-row w-full h-16 bg-[#FFFFFF42] border-t border-white justify-between items-center" aria-label="mobile bottom navigation">
                    <ul className="flex flex-row w-full justify-evenly">
                        <li>
                            <a href="">
                                <span className="sr-only">Home</span><HomeIcon />
                            </a>
                        </li>
                        <li>
                            <a href="">
                                <span className="sr-only">Home</span><HomeIcon />
                            </a>
                        </li>
                        <li>
                            <a href="">
                                <span className="sr-only">Home</span><HomeIcon />
                            </a>
                        </li>
                        <li>
                            <a href="">
                                <span className="sr-only">Home</span><HomeIcon />
                            </a>
                        </li>
                    </ul>
                </nav>
            </div>

        </>
    )
}

export default MobileBottomNav