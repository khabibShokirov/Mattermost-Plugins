// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';

import {Client4} from 'mattermost-redux/client';

import manifest from '@/manifest';

// Задай реальный bot id!
const BOT_USER_ID = 'alfabuddy_bot_userid';

const MyIcon = () => (
    <svg
        width={20}
        height={20}
        fill='currentColor'
    >
        <circle
            cx='10'
            cy='10'
            r='10'
            fill='#888'
        />
    </svg>
);

export default class Plugin {
    async initialize(registry: any, store: any) {
        if (typeof registry.registerChannelHeaderButtonAction === 'function') {
            registry.registerChannelHeaderButtonAction(
                <MyIcon/>,
                async () => {
                    const currentUserId = store.getState().entities.users.currentUserId;
                    const channel = await Client4.createDirectChannel([
                        currentUserId,
                        BOT_USER_ID,
                    ]);
                    if (channel && channel.name) {
                        window.location.href = `/channels/${channel.name}`;
                    }
                },
                'Открыть чат с Alfabuddy',
            );
        }
    }
}

declare global {
    interface Window {
        registerPlugin(pluginId: string, plugin: Plugin): void;
    }
}
window.registerPlugin(manifest.id, new Plugin());
