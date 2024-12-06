"use client";

import useQueryParams from "@/hooks/useQueryParams";
import { Event } from "@/types/group";
import React, { useEffect, useState } from "react";

export default function Events({ groupId }: { groupId: string }) {
    const { limit, offset } = useQueryParams();
    const [events, setEvents] = useState<Event[]>([]);

    const registerToEvent = (id: string) => async () =>
        await fetch(`/api/groups/${groupId}/events/${id}`, { method: "POST" });

    useEffect(() => {
        const fetchEvents = async () => {
            const response = await fetch(
                `/api/groups/${groupId}/events?limit=${limit}&offset=${offset}`
            );
            const data: Event[] = await response.json();

            setEvents(data);
        };

        fetchEvents();
    }, [groupId, limit, offset]);

    return (
        <>
            {events.length > 0 ? (
                <ul className="flex flex-col gap-2">
                    {events.map((event, idx) => (
                        <li key={idx}>
                            <h3>{event.title}</h3>
                            <p>{event.description}</p>
                            <div>{event.date}</div>
                            <label htmlFor="going">Going</label>
                            <input
                                type="checkbox"
                                name="going"
                                defaultChecked={event.going}
                                onChange={registerToEvent(event.id)}
                                id="going"
                            />
                        </li>
                    ))}
                </ul>
            ) : (
                <span>No events</span>
            )}
        </>
    );
}
