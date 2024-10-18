import Actualite from "@/components/Actualite";
import Chat from "@/components/Chat";
import Header from "@/components/Header";
import Layout from "@/components/Layout";
import SideNavMenu from "@/components/SideNavMenu";

export default function Home() {
    return <Layout>
        <Header />
        <main>
            <div className="hidden absolute left-0 top-[150px] xl:flex">
                <SideNavMenu />
            </div>
            <div className="absolute mt-3  left-1/2 -translate-x-1/2"><Actualite /></div>
            <div className="hidden absolute right-0 xl:flex"><Chat /></div>
        </main>
    </Layout >;
}
