import React from 'react';
import {
    useSortable,
    SortableContext,
    verticalListSortingStrategy
} from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';
import { Row, Col, Button, Badge } from 'antd';
import { PlusOutlined } from '@ant-design/icons';

// Column
export const SectionItem = (props) => {
    const { id, items, title, data, isSortingContainer, dragOverlay, onAdd } = props;
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
        let h = document.getElementsByClassName('kanban-column')[0].clientHeight;
        return h;
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
                    count={items.length ? items.length : 0}
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
export const FieldItem = (props) => {
    const { id, item, dragOverlay } = props;
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
        disabled: props.disabled,
        data: {
            type: 'FIELD'
        }
    });

    const style = {
        transform: CSS.Translate.toString(transform),
        transition,
        opacity: isDragging ? 0.5 : 1,
        boxShadow: dragOverlay
            ? '0 0 0 calc(1px / 1) rgba(63, 63, 68, 0.05), -1px 0 15px 0 rgba(34, 33, 81, 0.01), 0px 15px 15px 0 rgba(34, 33, 81, 0.25)'
            : '',
        border: dragOverlay ? '1px solid rgba(64, 150, 255, 1)' : '1px solid #dcdcdc', // 1px solid rgba(64, 150, 255, 1)
        cursor: dragOverlay ? 'grabbing' : 'grab',
        //transform: dragOverlay ? 'rotate(0deg) scale(1.02)' : 'rotate(0deg) scale(1.0)'
        touchAction:
            window.PointerEvent ||
            'ontouchstart' in window ||
            navigator.MaxTouchPoints > 0 ||
            navigator.msMaxTouchPoints > 0
                ? 'manipulation'
                : 'none'
    };
    return (
        <div
            ref={props.disabled ? null : setNodeRef}
            className="card"
            style={style}
            {...attributes}
            {...listeners}
        >
            <div>
                <Row justify="space-between">
                    <Col span={20}>{item.title}</Col>
                </Row>
            </div>
        </div>
    );
};
