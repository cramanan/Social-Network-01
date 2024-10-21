import React from 'react'

const ProfileSetting = () => {
    const profileData = [
        { label: "Name", value: "Name" },
        { label: "FirstName", value: "FirstName" },
        { label: "Date of birth", value: "Date of birth" },
        { label: "Nickname", value: "Nickname" },
        { label: "About me", value: <textarea name="about-me" id="about-me" rows={3} cols={20} className="resize-x-none"></textarea> },
    ];
    return (
        <>
            <div className="w-[700px] h-[400px] relative flex flex-col items-center bg-white bg-opacity-40 rounded-3xl">
                <div className="w-[400px] h-12 flex items-center justify-center border border-gray-400 rounded-3xl relative mt-5">
                    <h1 className="text-black text-xl font-semibold font-['Inter']">Profile</h1>
                </div>

                <div className="flex flex-row w-full items-center gap-20 pl-10 pr-5 mt-10">
                    <div className="w-[150px] h-[150px] bg-white rounded-full"></div>

                    <div>
                        <table className="w-full border-collapse">
                            <tbody>
                                {profileData.map((item, index) => (
                                    <tr key={index}>
                                        <td className="w-[200px] p-1 font-bold">{item.label}</td>
                                        <td className="w-[200px] p-1">{item.value}</td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        </>
    )
}

export default ProfileSetting