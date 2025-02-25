import React, { useState } from 'react';
import {
    useSortable,
    SortableContext,
    verticalListSortingStrategy
} from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';
import { Row, Col, Button, Badge, Space, Avatar } from 'antd';
import { PlusOutlined, UserOutlined } from '@ant-design/icons';

// Column
export const SectionItem = (props) => {
    const { id, items, title, data, isSortingContainer, dragOverlay, onAdd, onView } =
        props;
    const {
        //active,
        attributes,
        isDragging,
        listeners,
        //over,
        setNodeRef,
        setActivatorNodeRef,
        transition,
        transform
    } = useSortable({
        id: id,
        data: {
            type: 'SECTION'
        }
    });

    const getColumnHeight = () => {
        return document.getElementsByClassName('kanban-column')[0].clientHeight;
    };

    const style = {
        transform: CSS.Translate.toString(transform),
        height: dragOverlay ? `${getColumnHeight() + 'px'}` : null,
        transition,
        opacity: isDragging ? 0.5 : 1,
        boxShadow: dragOverlay
            ? '0 0 0 calc(1px / 1) rgba(63, 63, 68, 0.05), -1px 0 15px 0 rgba(34, 33, 81, 0.01), 0px 15px 15px 0 rgba(34, 33, 81, 0.25)'
            : '',
        border: dragOverlay ? '1px solid rgba(64, 150, 255, 1)' : '1px solid #dcdcdc', // 1px solid rgba(64, 150, 255, 1)
        //cursor: dragOverlay ? "grabbing" : "grab",
        //transform: dragOverlay ? 'rotate(0deg) scale(1.02)' : 'rotate(0deg) scale(1.0)'
        touchAction:
            'ontouchstart' in window ||
            navigator.MaxTouchPoints > 0 ||
            navigator.msMaxTouchPoints > 0
                ? 'manipulation'
                : 'none'
    };

    return (
        <div
            ref={setNodeRef}
            className="kanban-column"
            style={style}
            //{...attributes}
            //{...listeners}
        >
            <div
                ref={setActivatorNodeRef}
                {...attributes}
                {...listeners}
                className="kanban-column-header"
                style={{
                    cursor: dragOverlay ? 'grabbing' : 'grab'
                }}
            >
                {title}
                <Badge
                    count={items.length || 0}
                    showZero={true}
                    style={{
                        backgroundColor: '#fff',
                        color: '#000',
                        marginLeft: '10px'
                    }}
                />
            </div>
            <div className="kanban-column-list">
                <SortableContext
                    id={id}
                    items={items}
                    strategy={verticalListSortingStrategy} // verticalListSortingStrategy rectSortingStrategy
                >
                    {items.map((item, _index) => {
                        return (
                            <FieldItem
                                id={item}
                                key={item}
                                item={data.filter((d) => 'task-' + d.id === item)[0]}
                                disabled={isSortingContainer}
                                onView={onView}
                            />
                        );
                    })}
                </SortableContext>
            </div>
            <div className="kanban-column-footer">
                <Button
                    type="text"
                    icon={<PlusOutlined />}
                    size="small"
                    style={{ width: '100%', textAlign: 'left' }}
                    onClick={() => onAdd(parseInt(id.split('-')[1]))}
                >
                    Add task
                </Button>
            </div>
        </div>
    );
};

// Task
export const FieldItem = ({ id, item, dragOverlay, disabled, onView }) => {
    const [isItemDragging, setIsItemDragging] = useState(false);
    const {
        setNodeRef,
        //setActivatorNodeRef,
        listeners,
        isDragging,
        //isSorting,
        //over,
        //overIndex,
        transform,
        transition,
        attributes
    } = useSortable({
        id: id,
        disabled,
        data: {
            type: 'FIELD'
        }
    });

    const handleMouseDown = () => {
        setIsItemDragging(false);
    };

    const handleMouseMove = () => {
        setIsItemDragging(true);
    };

    const handleMouseUp = (id) => {
        if (isItemDragging) {
            return;
        }
        onView(id);
    };

    const style = {
        transform: CSS.Translate.toString(transform),
        transition,
        opacity: isDragging ? 0.5 : 1,
        boxShadow: dragOverlay
            ? '0 0 0 calc(1px / 1) rgba(63, 63, 68, 0.05), -1px 0 15px 0 rgba(34, 33, 81, 0.01), 0px 15px 15px 0 rgba(34, 33, 81, 0.25)'
            : '',
        // border: dragOverlay ? '1px solid rgba(64, 150, 255, 1)' : '1px solid #dcdcdc',
        cursor: dragOverlay ? 'grabbing' : 'grab',
        //transform: dragOverlay ? 'rotate(0deg) scale(1.02)' : 'rotate(0deg) scale(1.0)'
        padding: '0',
        touchAction:
            window.PointerEvent ||
            'ontouchstart' in window ||
            navigator.MaxTouchPoints > 0 ||
            navigator.msMaxTouchPoints > 0
                ? 'manipulation'
                : 'none'
    };
    /*
    if (item.color) {
        style.borderColor = item.color;
    }
    */
    return (
        <div
            ref={disabled ? null : setNodeRef}
            className="card"
            style={style}
            {...attributes}
            {...listeners}
        >
            <div
                onMouseDown={handleMouseDown}
                onMouseMove={handleMouseMove}
                onMouseUp={() => {
                    handleMouseUp(item.id);
                }}
                style={{
                    width: '100%',
                    padding: '10px'
                }}
            >
                <Row justify="space-between" className="task-item">
                    <Col span={24}>
                        <Badge size="small" count={0} offset={[9, -14]}>
                            <span>{item.title}</span>
                        </Badge>
                    </Col>
                </Row>

                <Row
                    justify="space-between"
                    style={{
                        marginTop: '10px',
                        color: '#777'
                    }}
                >
                    <Col>
                        <Space align="center">
                            {item.task_users.length > 0 && (
                                <Avatar.Group max={{ count: 2 }} size="small">
                                    {item.task_users.map((user) => {
                                        return (
                                            <Avatar
                                                icon={<UserOutlined />}
                                                src={user.avatar}
                                                key={user.id}
                                            />
                                        );
                                    })}
                                </Avatar.Group>
                            )}
                        </Space>
                    </Col>
                </Row>
            </div>
        </div>
    );
};
