"use client";

import useQueryParams from "@/hooks/useQueryParams";
import { Event } from "@/types/group";
import React, { useEffect, useState } from "react";

export default function Events({ groupId }: { groupId: string }) {
    const { limit, offset } = useQueryParams();
    const [events, setEvents] = useState<Event[]>([]);

    const changeGoingState = (eventId: string) => async () => {
        /*const response =*/ await fetch(
            `/api/group/${groupId}/events/${eventId}`,
            { method: "POST" }
        );
    };

    useEffect(() => {
        const fetchEvents = async () => {
            const response = await fetch(
                `/api/group/${groupId}/events?limit=${limit}&offset=${offset}`
            );
            const data: Event[] = await response.json();

            setEvents(data);
        };

        fetchEvents();
    }, [groupId, limit, offset]);

    return (
        <ul>
            {events.map((event, idx) => (
                <li key={idx}>
                    <h1>{event.title}</h1>
                    <p>{event.description}</p>
                    <div>{event.date}</div>
                    <label htmlFor="going">Going</label>
                    <input
                        type="checkbox"
                        name="going"
                        defaultChecked={event.going}
                        onChange={changeGoingState(event.id)}
                        id="going"
                    />
                </li>
            ))}
        </ul>
    );
}
