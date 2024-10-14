import React from 'react'

const ProfileStats = () => {
    const items = [
        { label: "Publication(s)", value: 5 },
        { label: "Follower(s)", value: 5 },
        { label: "Follow(s)", value: 5 },
        { label: "Like(s)", value: 5 }
    ]

    return (
        <div className='w-[400px] h-16 bg-white/30 rounded-2xl flex flex-row items-center justify-between px-3'>
            {items.map((item, index) => (
                <div key={index} className="flex flex-col items-center w-[86px]">
                    <div className="font-bold">{item.value}</div>
                    <div className="text-black/50">{item.label}</div>
                </div>
            ))}
            {/* <div className={divStyle}>
                <div className={statStyle}>5</div>
                <div className={labelStyle}>Publications</div>
            </div>
            <div className={divStyle}>
                <div className={statStyle}>5</div>
                <div className={labelStyle}>Followers</div>
            </div>
            <div className={divStyle}>
                <div className={statStyle}>5</div>
                <div className={labelStyle}>Follows</div>
            </div>
            <div className={divStyle}>
                <div className={statStyle}>5</div>
                <div className={labelStyle}>Like</div>
            </div> */}
        </div>
    )
}

export default ProfileStats