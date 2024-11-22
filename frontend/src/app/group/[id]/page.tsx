
import React from "react";
import { Group } from "@/types/group";
import { Params } from "@/types/query";
import NewEvent from "./NewEvent";
import Events from "./Events";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { BackIcon } from "@/components/icons/BackIcon";
import Link from "next/link";
import Image from "next/image";

export default async function GroupPage({ params }: { params: Params }) {
    const { id } = await params;
    // const [showMemberList, setShowMemberList] = useState(true)
    // const [showEventList, setShowEventList] = useState(false)

    const response = await fetch(
        `${process.env.NEXT_PUBLIC_API_URL}/api/group/${id}`
    );
    const group: Group = await response.json();

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
                    <div className="flex flex-row justify-between items-center w-full h-16 px-5 shadow-xl">
                        <Link href={"/group"}>
                            <BackIcon />
                        </Link>


                        <div className="flex flex-col justify-center items-center">
                            <h1 className="font-bold">{group.name}</h1>
                            <p>{group.description}</p>
                        </div>

                        <Image src={group.image} alt="" width={50} height={50} />
                    </div>

                    <div className="flex flex-row w-full h-full">
                        <div className="flex flex-col items-center w-72 border-r-4">
                            <NewEvent groupId={group.id} />

                            <span className="font-bold">Members</span>

                            <span>List of Members</span>

                            <span className="font-bold">Events</span>
                            <Events groupId={group.id} />
                        </div>

                        <div className="w-full">Group posts</div>
                    </div>
                </div>
            </HomeProfileLayout >
        </>
    );
}
