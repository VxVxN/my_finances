'use client';

import React, {FC} from 'react';
import {Modal, ModalBody, ModalContent, ModalFooter, ModalHeader, useDisclosure} from '@nextui-org/modal';
import {Button} from '@nextui-org/react';
import {Input} from '@nextui-org/input';

interface CreateOrderFormProps {
    isOpen: boolean;
    onOpenChange: (event: boolean) => void;
}

export const CreateOrderForm: FC<CreateOrderFormProps> = ({isOpen, onOpenChange}) => {
    return (
        <Modal
            isOpen={isOpen}
            placement="auto"
            onOpenChange={onOpenChange}
        >
            <ModalContent>
                {(onClose) => (
                    <>
                        <ModalHeader className="flex flex-col gap-1">Добавить сделку</ModalHeader>
                        <ModalBody>
                            <Input isRequired type="text" label="Название монеты" placeholder='Введите монету' labelPlacement="outside"/>
                            <Input
                                isRequired
                                type="number"
                                label="Цена"
                                placeholder="0.00"
                                labelPlacement="outside"
                                startContent={
                                    <div className="pointer-events-none flex items-center">
                                        <span className="text-default-400 text-small">$</span>
                                    </div>
                                }
                            />
                            <Input isRequired type="number" placeholder='0' label="Колличество" labelPlacement="outside"/>
                        </ModalBody>
                        <ModalFooter>
                            <Button color="danger" variant="light" onPress={onClose}>
                                Закрыть
                            </Button>
                            <Button color="primary" onPress={onClose}>
                                Создать
                            </Button>
                        </ModalFooter>
                    </>
                )}
            </ModalContent>
        </Modal>
    );
};