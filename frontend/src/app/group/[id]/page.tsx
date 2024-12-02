import React from "react";
import { Group } from "@/types/group";
import { Params } from "@/types/query";
import NewEvent from "./NewEvent";
import Events from "./Events";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { BackIcon } from "@/components/icons/BackIcon";
import Link from "next/link";
import Image from "next/image";
import { NewPost } from "@/components/NewPost";
import { Post } from "@/types/post";
import PostComponent from "@/components/PostComponent";

export default async function GroupPage({ params }: { params: Params }) {
    const { id } = await params;

    // const [showMemberList, setShowMemberList] = useState(true)
    // const [showEventList, setShowEventList] = useState(false)

    const response = await fetch(
        `http://${process.env.NEXT_PUBLIC_API_URL}/api/group/${id}`
    );
    const group: Group = await response.json();

    // const { limit, offset, next, previous } = useQueryParams();

    const test = await fetch(
        `http://${process.env.NEXT_PUBLIC_API_URL}/api/group/${id}/posts?limit=20&offset=0`
    );
    const posts: Post[] = await test.json();

    // const handleMemberListClick = () => {
    //     setShowMemberList(showMemberList)
    //     setShowEventList(false)
    // }

    // const handleEventListClick = () => {
    //     setShowMemberList(false)
    //     setShowEventList(showEventList)
    // }

    return (
        <>
            <HomeProfileLayout>
                <div className="flex flex-col w-screen h-[calc(100vh-185px)] xl:bg-white/25 z-10 xl:mt-3 xl:w-[900px] lg:rounded-t-[25px] xl:h-[calc(100vh-70px)]">
                    <div className="flex flex-row justify-between items-center w-full h-16 px-5 py-2 shadow-xl">
                        <Link href={"/group"}>
                            <BackIcon />
                        </Link>

                        <div className="flex flex-col justify-center items-center">
                            <h1 className="font-bold">{group.name}</h1>
                            <p>{group.description}</p>
                        </div>

                        <div className="flex flex-row gap-5 text-3xl">
                            <Image
                                src={group.image}
                                alt=""
                                width={50}
                                height={50}
                            />

                            {/* displaying only if in group */}
                            {/* Invite followers to group */}
                            <input type="button" value="chat" />
                            <input type="button" value="+" className="font-bold" />
                        </div>
                    </div>


                    {/* Display if not in group */}
                    {/* <div className="flex flex-col items-center font-bold text-3xl gap-5">
                        <h2>You are not in the group yet, <br /> click below to send a request !</h2>

                        <label htmlFor="request-to-group"></label>
                        <input name="request-to-group" id="request-to-group" type="button" value="request to join" />
                    </div> */}


                    {/* Display if in group */}
                    <div className="flex flex-row w-full h-full">
                        <div className="flex flex-col items-center w-72 border-r-4">
                            <div className="flex flex-col pt-3 gap-2">
                                <NewPost groupId={id} />
                                <NewEvent groupId={group.id} />
                            </div>

                            <ul className="flex flex-col items-center">
                                <li className="font-bold">Invites</li>
                                <span>List of Invites</span>

                                <li className="font-bold">Members</li>
                                <span>List of Members</span>

                                <li className="font-bold">Events</li>
                                <Events groupId={group.id} />
                            </ul>
                        </div>

                        <div className="flex flex-col w-full p-3 gap-3 overflow-scroll no-scrollbar">
                            {posts.map((post, idx) =>
                                <PostComponent key={idx} post={post} />
                            )}
                        </div>
                    </div>
                </div>
            </HomeProfileLayout>
        </>
    );
}
