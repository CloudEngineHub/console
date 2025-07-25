/**
 * Copyright 2022 Redpanda Data, Inc.
 *
 * Use of this software is governed by the Business Source License
 * included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
 *
 * As of the Change Date specified in that file, in accordance with
 * the Business Source License, use of this software will be governed
 * by the Apache License, Version 2.0
 */

import { Badge, Box, Breadcrumbs, Button, ColorModeSwitch, CopyButton, Flex, Text } from '@redpanda-data/ui';
import { computed } from 'mobx';
import { observer } from 'mobx-react';
import { Link as ReactRouterLink, useMatch } from 'react-router-dom';
import { isEmbedded } from '../../config';
import { api } from '../../state/backendApi';
import { type BreadcrumbEntry, uiState } from '../../state/uiState';
import { IsDev } from '../../utils/env';
import DataRefreshButton from '../misc/buttons/data-refresh/Component';
import { UserPreferencesButton } from '../misc/UserPreferences';

const AppPageHeader = observer(() => {
  const showRefresh = useShouldShowRefresh();
  const showBetaBadge = useShouldShowBetaBadge();

  const breadcrumbItems = computed(() => {
    const items: BreadcrumbEntry[] = [...uiState.pageBreadcrumbs];

    if (!isEmbedded() && uiState.selectedClusterName) {
      items.unshift({
        heading: '',
        title: 'Cluster',
        linkTo: '/',
      });
    }

    return items;
  }).get();

  const lastBreadcrumb = breadcrumbItems.pop();

  return (
    <Box>
      {/* we need to refactor out #mainLayout > div rule, for now I've added this box as a workaround */}
      <Flex mb={5} alignItems="center" justifyContent="space-between">
        {!isEmbedded() && (
          <Breadcrumbs
            showHomeIcon={false}
            items={breadcrumbItems.map((x) => ({
              name: x.title,
              heading: x.heading,
              to: x.linkTo,
            }))}
          />
        )}
      </Flex>

      <Flex pb={2} alignItems="center" justifyContent="space-between">
        <Flex alignItems="center">
          {lastBreadcrumb && (
            <Text
              fontWeight={700}
              as="span"
              fontSize="xl"
              mr={2}
              {...(lastBreadcrumb.options?.canBeTruncated
                ? {
                    wordBreak: 'break-all',
                    whiteSpace: 'break-spaces',
                  }
                : {
                    whiteSpace: 'nowrap',
                  })}
            >
              {lastBreadcrumb.title}
            </Text>
          )}
          {showBetaBadge && <Badge ml={2}>beta</Badge>}
          {lastBreadcrumb && (
            <Box>
              {lastBreadcrumb.options?.canBeCopied && <CopyButton content={lastBreadcrumb.title} variant="ghost" />}
            </Box>
          )}
          {showRefresh && <DataRefreshButton />}
        </Flex>
        <Flex alignItems="center" gap={2}>
          {!isEmbedded() && api.isRedpanda && (
            <Button
              as={ReactRouterLink}
              to={api.userData?.canViewDebugBundle ? '/debug-bundle' : undefined}
              variant="ghost"
              isDisabled={!api.userData?.canViewDebugBundle}
              tooltip={
                !api.userData?.canViewDebugBundle ? 'You need RedpandaCapability.MANAGE_DEBUG_BUNDLE permission' : null
              }
            >
              Debug bundle
            </Button>
          )}
          <UserPreferencesButton />
          {IsDev && !isEmbedded() && <ColorModeSwitch m={0} p={0} variant="ghost" />}
        </Flex>
      </Flex>
    </Box>
  );
});

export default AppPageHeader;

/**
 * Custom React Hook: Determines whether to show the refresh button based on route matches.
 * It checks various routes and conditions to decide if the refresh button should be displayed
 * in the header next to the breadcrumb.
 *
 * @returns {boolean} Indicates whether the refresh button should be shown (true/false).
 */
function useShouldShowRefresh() {
  const connectClusterMatch = useMatch({
    path: '/connect-clusters/:clusterName/:connectorName',
    end: false,
  });

  const schemaCreateMatch = useMatch({
    path: '/schema-registry/create',
    end: false,
  });

  const topicProduceRecordMatch = useMatch({
    path: '/topics/:topicName/produce-record',
    end: false,
  });

  const secretsMatch = useMatch({
    path: '/secrets',
    end: true,
  });

  const agentsMatch = useMatch({
    path: '/agents',
    end: true,
  });

  const agentDetailsMatch = useMatch({
    path: '/agents/:agentId',
    end: true,
  });

  const createAgentMatch = useMatch({
    path: '/agents/create',
    end: false,
  });

  if (connectClusterMatch && connectClusterMatch.params.connectorName === 'create-connector') return false;
  if (schemaCreateMatch) return false;
  if (topicProduceRecordMatch) return false;
  if (secretsMatch) return false;
  if (agentsMatch) return false;
  if (agentDetailsMatch) return false;
  if (createAgentMatch) return false;

  return true;
}

function useShouldShowBetaBadge() {
  const agentsMatch = useMatch({
    path: '/agents',
    end: false,
  });

  return agentsMatch !== null;
}
