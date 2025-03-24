import { Button, ButtonGroup, Flex, Spinner, Stack, Tabs, Text, useDisclosure } from '@redpanda-data/ui';
import type { TabsItemProps } from '@redpanda-data/ui/dist/components/Tabs/Tabs';
import { runInAction } from 'mobx';
import type { Pipeline } from 'protogen/redpanda/api/dataplane/v1/pipeline_pb';
import { useEffect } from 'react';
import { type Agent, useGetAgentQuery } from 'react-query/api/agent';
import { useParams } from 'react-router-dom';
import { uiState } from 'state/uiState';
import { capitalizeFirst } from 'utils/utils';
import { DeleteAgentModal } from '../delete-agent-modal';
import { AgentPipelineTab } from './agent-pipeline-tab';
import { AgentChatTab } from './chat/agent-chat-tab';

// Hack for MobX to ensure we don't need to use observables
export const updatePageTitle = ({ agent }: { agent: Agent | undefined }) => {
  const nameWithoutPrefix = agent?.displayName ?? '';
  runInAction(() => {
    uiState.pageTitle = `Agent ${nameWithoutPrefix}`;
    uiState.pageBreadcrumbs.pop(); // Remove last breadcrumb to ensure the agent title is used without previous page breadcrumb being shown
    uiState.pageBreadcrumbs.push({
      title: `Agent ${nameWithoutPrefix}`,
      linkTo: `/agents/${agent?.id}`,
      heading: 'Agent Details',
    });
  });
};

export const AgentDetailsPage = () => {
  const { agentId } = useParams<{ agentId: Pipeline['id'] }>();
  const { data: agentData, isLoading: isAgentDataLoading } = useGetAgentQuery({ id: agentId });

  const {
    isOpen: isDeleteAgentModalOpen,
    onOpen: onDeleteAgentModalOpen,
    onClose: onDeleteAgentModalClose,
  } = useDisclosure();

  useEffect(() => {
    updatePageTitle({ agent: agentData?.agent });
  }, [agentData?.agent]);

  if (isAgentDataLoading) {
    return (
      <Flex justifyContent="center" padding={8}>
        <Spinner size="lg" />
      </Flex>
    );
  }

  const tabs: TabsItemProps[] = [
    {
      key: 'chat',
      name: 'Chat',
      component: (
        <AgentChatTab
          agent={agentData?.agent?.pipelines?.find(
            (pipeline) => pipeline?.tags?.__redpanda_cloud_pipeline_purpose === 'chat',
          )}
        />
      ),
    },
    ...(agentData?.agent?.pipelines ?? []).map((pipeline) => ({
      key: pipeline?.id ?? '',
      name: capitalizeFirst(pipeline?.displayName ?? ''),
      component: <AgentPipelineTab pipeline={pipeline} />,
    })),
  ];

  return (
    <>
      <Stack spacing={8}>
        <Stack spacing={4}>
          <Text>{agentData?.agent?.description}</Text>
          <ButtonGroup>
            <Button
              variant="outline-delete"
              onClick={() => {
                onDeleteAgentModalOpen();
              }}
              data-testid="delete-agent-button"
            >
              Delete
            </Button>
          </ButtonGroup>
        </Stack>
        <Tabs items={tabs} />
      </Stack>
      <DeleteAgentModal isOpen={isDeleteAgentModalOpen} onClose={onDeleteAgentModalClose} agent={agentData?.agent} />
    </>
  );
};
