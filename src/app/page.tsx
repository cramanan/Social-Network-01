import Actualite from "@/components/Actualite";
import Chat from "@/components/Chat";
import ChatBox from "@/components/ChatBox";
import ChatList from "@/components/ChatList";
import Comment from "@/components/Comment";
import CreatePost from "@/components/CreatePost";
import FindUser from "@/components/FindUser";
import FriendInvite from "@/components/FriendInvite";
import FriendInviteList from "@/components/FriendInviteList";
import Header from "@/components/Header";
import { ProfileCircle } from "@/components/icons/ProfileCircle";
import Layout from "@/components/Layout";
import Media from "@/components/Media";
import Post from "@/components/Post";
import ProfileBanner from "@/components/ProfileBanner";
import ProfilePost from "@/components/ProfilePost";
import ProfileSetting from "@/components/ProfileSetting";
import ProfileStats from "@/components/ProfileStats";
import Publication from "@/components/Publication";
import SideNavMenu from "@/components/SideNavMenu";
import UserList from "@/components/UserList";
import Users from "@/components/Users";

export default function Home() {
    return <Layout>
        <Header />
        {/* <main>
            <div className="absolute left-0 top-[150px]">
                <SideNavMenu />
            </div>
            <div className="absolute left-1/2 -translate-x-1/2"><Actualite /></div>
            <div className="absolute right-0"><Chat /></div>
        </main> */}
        <main className="flex flex-col gap-5 justify-center items-center">
            <ProfileSetting />
            {/* <Actualite />
            <Users />
            <Chat />
            <ChatList />
            <FindUser />
            <UserList />
            <ChatBox />
            <Comment />
            <CreatePost />
            <Media />
            <Post />
            <ProfileBanner />
            <ProfileStats />
            <ProfilePost />
            <Publication />
            <SideNavMenu />
            <FriendInvite />
            <FriendInviteList /> */}
        </main>
    </Layout>;
}
