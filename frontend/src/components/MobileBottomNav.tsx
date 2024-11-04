'use client'

import { HomeIcon } from "./icons/HomeIcon"
import { FindUserIcon } from "./icons/FindUserIcon"
import { UserListIcon } from "./icons/UserListIcon"
import ChatIcon from "./icons/ChatIcon"

const MobileBottomNav = () => {
    return (
        <>
            <div className="relative w-full h-full z-60">
                <nav className="relative flex flex-row w-full h-16 bg-[#FFFFFF42] border-t border-white justify-between items-center" aria-label="mobile bottom navigation">
                    <ul className="flex flex-row w-full justify-evenly">
                        <li>
                            <a href="/">
                                <span className="sr-only">Home</span><HomeIcon />
                            </a>
                        </li>
                        <li>
                            <button>
                                <span className="sr-only">UserList</span><UserListIcon />
                            </button>
                        </li>
                        <li>
                            <button>
                                <span className="sr-only">FindUser</span><FindUserIcon />
                            </button>
                        </li>
                        <li>
                            <button>
                                <span className="sr-only">Chat</span><ChatIcon />
                            </button>
                        </li>
                    </ul>
                </nav>
            </div>
        </>
    )
}

export default MobileBottomNav